package snipe

import (
	"encoding/json"
	"net/http"
)

type SnipeRespBody struct {
	response *http.Response
	body     []byte
	errMsg   error
}

type SnipeResponse struct {
	Total int        `json:"total,omitempty"`
	Rows  []Hardware `json:"rows,omitempty"`
}

type Hardware struct {
	ID               int              `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	AssetTag         string           `json:"asset_tag,omitempty"`
	Serial           string           `json:"serial,omitempty"`
	Model            Model            `json:"model,omitempty"`
	ModelNumber      string           `json:"model_number,omitempty"`
	Eol              Eol              `json:"eol,omitempty"`
	StatusLabel      StatusLabel      `json:"status_label,omitempty"`
	Category         Category         `json:"category,omitempty"`
	Manufacturer     Manufacturer     `json:"manufacturer,omitempty"`
	Supplier         Supplier         `json:"supplier,omitempty"`
	Notes            string           `json:"notes,omitempty"`
	OrderNumber      string           `json:"order_number,omitempty"`
	Company          interface{}      `json:"company,omitempty"`
	Location         Location         `json:"location,omitempty"`
	RtdLocation      RtdLocation      `json:"rtd_location,omitempty"`
	Image            string           `json:"image,omitempty"`
	Qr               string           `json:"qr,omitempty"`
	AltBarcode       interface{}      `json:"alt_barcode,omitempty"`
	AssignedTo       AssignedTo       `json:"assigned_to,omitempty"`
	WarrantyMonths   interface{}      `json:"warranty_months,omitempty"`
	WarrantyExpires  interface{}      `json:"warranty_expires,omitempty"`
	CreatedAt        CreatedAt        `json:"created_at,omitempty"`
	UpdatedAt        UpdatedAt        `json:"updated_at,omitempty"`
	LastAuditDate    interface{}      `json:"last_audit_date,omitempty"`
	NextAuditDate    interface{}      `json:"next_audit_date,omitempty"`
	DeletedAt        interface{}      `json:"deleted_at,omitempty"`
	PurchaseDate     PurchaseDate     `json:"purchase_date,omitempty"`
	Age              string           `json:"age,omitempty"`
	LastCheckout     interface{}      `json:"last_checkout,omitempty"`
	ExpectedCheckin  interface{}      `json:"expected_checkin,omitempty"`
	PurchaseCost     string           `json:"purchase_cost,omitempty"`
	CheckinCounter   int              `json:"checkin_counter,omitempty"`
	CheckoutCounter  int              `json:"checkout_counter,omitempty"`
	RequestsCounter  int              `json:"requests_counter,omitempty"`
	UserCanCheckout  bool             `json:"user_can_checkout,omitempty"`
	CustomFields     CustomFields     `json:"custom_fields,omitempty"`
	AvailableActions AvailableActions `json:"available_actions,omitempty"`
}

type AssignedTo struct {
	ID             int    `json:"id,omitempty"`
	UserName       string `json:"username,omitempty"`
	Name           string `json:"name,omitempty"`
	FirstName      string `json:"first_name,omitempty"`
	LastName       string `json:"last_name,omitempty"`
	EmployeeNumber int    `json:"employee_number,omitempty"`
	Type           string `json:"type,omitempty"`
}

type Model struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Eol struct {
	Date      string `json:"date,omitempty"`
	Formatted string `json:"formatted,omitempty"`
}

type StatusLabel struct {
	ID         int    `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	StatusType string `json:"status_type,omitempty"`
	StatusMeta string `json:"status_meta,omitempty"`
}

type Category struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Manufacturer struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Supplier struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Location struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type RtdLocation struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CreatedAt struct {
	Datetime  string `json:"datetime,omitempty"`
	Formatted string `json:"formatted,omitempty"`
}

type UpdatedAt struct {
	Datetime  string `json:"datetime,omitempty"`
	Formatted string `json:"formatted,omitempty"`
}

type PurchaseDate struct {
	Date      string `json:"date,omitempty"`
	Formatted string `json:"formatted,omitempty"`
}

type CustomFields struct {
}

type AvailableActions struct {
	Checkout bool `json:"checkout,omitempty"`
	Checkin  bool `json:"checkin,omitempty"`
	Clone    bool `json:"clone,omitempty"`
	Restore  bool `json:"restore,omitempty"`
	Update   bool `json:"update,omitempty"`
	Delete   bool `json:"delete,omitempty"`
}

type HardwareQueryOpts struct {
	Limit          int    `url:"limit,omitempty"`
	Offset         int    `url:"offset,omitempty"`
	Search         string `url:"search,omitempty"`
	OrderNumber    string `url:"order_number,omitempty"`
	Sort           string `url:"sort,omitempty"`
	Order          string `url:"order,omitempty"`
	ModelID        int    `url:"model_id,omitempty"`
	CategoryID     int    `url:"category_id,omitempty"`
	ManufacturerID int    `url:"manufacturer_id,omitempty"`
	CompanyID      int    `url:"company_id,omitempty"`
	LocationID     int    `url:"location_id,omitempty"`
	Status         string `url:"status,omitempty"`
	StatusID       string `url:"status_id,omitempty"`
}

type offsetRange struct {
	Limit  int
	Offset int
}

// GetSnipeAssets returns a full listing of assets in Snipe
// https://snipe-it.readme.io/reference/hardware-list
func (c *SnipeClient) GetSnipeAssets(opts *HardwareQueryOpts) ([]Hardware, int) {
	var snipeRes SnipeResponse
	req := c.NewRequest("GET", "hardware", nil)

	err := json.Unmarshal(req.body, &snipeRes)
	if err != nil {
		return nil, 0
	}

	return snipeRes.Rows, snipeRes.Total
}

func offSetter(total, limit, offset int) []offsetRange {
	o := []offsetRange{}

	keepGoing := true
	for keepGoing {
		o = append(o, offsetRange{
			Limit:  limit,
			Offset: offset,
		})

		if (offset + limit) <= total {
			offset += limit
		} else if (offset + limit) >= total {
			offset = ((total - offset) + offset)
		}

		if offset == total {
			break
		}
	}

	return o
}

func (c *SnipeClient) GetAllHardware(opts *HardwareQueryOpts) []Hardware {
	results, total := c.GetSnipeAssets(opts)

	if total > 500 {
		or := offSetter(total, opts.Limit, opts.Offset)
		for _, o := range or {
			opts.Limit = o.Limit
			opts.Offset = o.Offset

			hw, _ := c.GetSnipeAssets(opts)
			results = append(results, hw...)
		}
	}

	return results
}
