package analytics

type Sale struct {
	Product          string
	Category         string
	Quantity         int
	UnitPriceKopecks int64
	Cancelled        bool
}

type Report struct {
	RevenueKopecks    int64
	UnitsSold         int
	RevenueByCategory map[string]int64
	TopProduct        string
	InvalidSales      int
}

func AnalyzeSales(sales []Sale) Report {
	report := Report{}
	report.RevenueByCategory = make(map[string]int64)
	if len(sales) == 0 {
		return report
	}
	productQuantities := make(map[string]int)
	for _, sale := range sales {
		if sale.Cancelled {
			continue
		}
		if len(sale.Product) == 0 || len(sale.Category) == 0 || sale.Quantity <= 0 || sale.UnitPriceKopecks <= 0 {
			report.InvalidSales++
			continue
		}
		saleRevenue := int64(sale.Quantity) * sale.UnitPriceKopecks
		report.RevenueKopecks += saleRevenue
		report.UnitsSold += sale.Quantity
		report.RevenueByCategory[sale.Category] += saleRevenue
		productQuantities[sale.Product] += sale.Quantity
	}

	var maxUnits int
	for product, qty := range productQuantities {
		if report.TopProduct == "" || qty > maxUnits || (qty == maxUnits && product < report.TopProduct) {
			maxUnits = qty
			report.TopProduct = product
		}
	}
	return report
}
