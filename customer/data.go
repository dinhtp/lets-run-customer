package customer

const (
    Shopify     = "shopify"
    WooCommerce = "woocommerce"
)

type Platform struct {
    Shopify     string
    WooCommerce string
}

func (p *Platform) GetEndpoint(platformType string) string {
    switch platformType {
    case Shopify:
        return p.Shopify
    case WooCommerce:
        return p.WooCommerce
    default:
        return ""
    }
}
