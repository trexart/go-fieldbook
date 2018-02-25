# go-fieldbook

Fieldbook API for Go

## Requirements

To use this client, you'll need a key, secret and book id in Fieldbook.

TODO: Instructions on how to get these values

## Usage

### Create a Fieldbook Client

To create a fieldbook client and use it, you must import the library, and then create a new client using your API key, secret and book id.

    import (
        fieldbook "github.com/trexart/go-fieldbook"
    )

Create a client

    client := fieldbook.NewClient(KEY, SECRET, BOOK_ID)

### Struct for Sheet object

You'll need to create a struct matching the data in your sheet and map the fields to the json.

When sending an item to update, the Fieldbook API uses PATCH, so it will ignore everything that isn't sent. Because of that I've set all elements to 'omitempty'.

Your json fieldname will be the same as the title in your sheet, but lowercase and spaces replaced with underscores. For example, in the object below, I mapped the sheet heading of 'Product Name' to be Name. Therefore my json fieldname is 'product_name'. 

    type Product struct {
	    ID         int                `json:"id,omitempty"`
	    ImageURL   string             `json:"image,omitempty"`
	    SKU        string             `json:"sku,omitempty"`
	    Name       string             `json:"product_name,omitempty"`
	    Categories []Category         `json:"categories,omitempty"`
	    Price      float64            `json:"price,omitempty"`
	    Active     *bool               `json:"is_active,omitempty"`
    }

#### Booleans

In Go, if a boolean is set to false, it will not show up in the JSON. Therefore if you need to update a boolean value, you must make it a pointer, that way true or false always gets sent if it is set.

#### ID fields

If you don't do omitempty with an int, the JSON object sent to Fieldbook in an Insert will have ID of 0 and will give an error. You might need to have different objects for receiving and sending if you need these values to be better validated

#### Linked fields

Linked fields will always be returned as an array of values.

### Example queries

Get a list of items without any options

    var products []Product
    err := h.client.ListRecords("products", &products, nil)

List of items with options

    var products []Product
    options := fieldbook.QueryOptions{
		Expand: []string{"categories"},
	}
    err := h.client.ListRecords("products", &products, &options)

Get item

    var product Product
    err := h.client.GetRecord("products", id, &product, nil)

Create

    err = h.client.CreateRecord("products", product)

Update

    err = h.client.UpdateRecord("products", product.ID, product)