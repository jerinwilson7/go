package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	area() float64
	perim() float64
}

type Rectangle struct {
	height, width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (r Rectangle) perim() float64 {
	return 2*r.height + 2*r.width
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("Area is : ", g.area())
	fmt.Println("Perimeter is  :", g.perim())
}

func detectCircle(g Geometry) {
	if c, ok := g.(Circle); ok {
		fmt.Println("Circle with radius", c.radius)
	}
}

func detectRectangle(g Geometry) {
	if r, ok := g.(Rectangle); ok {
		fmt.Println("Rectangle with height and width of", r.height, r.width)
	}
}

func main() {

	circle1 := Circle{radius: 5}
	rectangle1 := Rectangle{height: 21, width: 25}

	measure(rectangle1)
	measure(circle1)

	detectCircle(circle1)
	detectCircle(rectangle1)

	detectRectangle(rectangle1)
	detectRectangle(circle1)

}


{
    "code": 0,
    "requests": {
        "request_status": "draft",
        "notes": "",
        "reminder_period": 1,
        "owner_id": "91891000000039003",
        "description": "",
        "request_name": "JERIN SIGN AGREEMENT",
        "modified_time": 1749183580760,
        "is_deleted": false,
        "expiration_days": 1,
        "is_sequential": false,
        "templates_used": [],
        "owner_first_name": "Vonnue",
        "sign_percentage": 0,
        "owner_email": "vonnue@pauzible.com",
        "created_time": 1749183580760,
        "email_reminders": false,
        "document_ids": [
            {
                "image_string": "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",
                "document_name": "deal_term",
                "pages": [
                    {
                        "image_string": "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",
                        "page": 0,
                        "is_thumbnail": true
                    }
                ],
                "document_size": 213162,
                "document_order": "0",
                "is_nom151_present": false,
                "is_editable": false,
                "total_pages": 3,
                "document_id": "91891000000042050"
            }
        ],
        "self_sign": false,
        "document_fields": [
            {
                "document_id": "91891000000042050",
                "fields": []
            }
        ],
        "in_process": false,
        "validity": -1,
        "request_type_name": "Others",
        "visible_sign_settings": {
            "allow_custom_signer_reason_visible_sign": true,
            "allow_predefined_reason_visible_sign": true,
            "visible_sign": false,
            "predefined_reasons_visible_sign": [
                "I authored this document",
                "I approve this document",
                "I reviewed and agree to this document"
            ],
            "allow_reason_visible_sign": true
        },
        "request_id": "91891000000042049",
        "request_type_id": "91891000000000127",
        "owner_last_name": "",
        "actions": [
            {
                "verify_recipient": false,
                "recipient_countrycode_iso": "",
                "action_type": "SIGN",
                "private_notes": "",
                "cloud_provider_name": "Zoho Sign",
                "recipient_email": "jerinwilson7@gmail.com",
                "send_completed_document": false,
                "recipient_phonenumber": "",
                "is_bulk": false,
                "action_id": "91891000000042066",
                "signing_order": 1,
                "cloud_provider_id": 10,
                "recipient_name": "jerin",
                "fields": [],
                "delivery_mode": "EMAIL",
                "action_status": "NOACTION",
                "is_signing_group": false,
                "recipient_countrycode": ""
            }
        ]
    },
    "message": "Document has been added",
    "status": "success"
}