// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package apigee

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceApigeeEnvironmentKeyvaluemapsEntries() *schema.Resource {
	return &schema.Resource{
		Create: resourceApigeeEnvironmentKeyvaluemapsEntriesCreate,
		Read:   resourceApigeeEnvironmentKeyvaluemapsEntriesRead,
		Delete: resourceApigeeEnvironmentKeyvaluemapsEntriesDelete,

		Importer: &schema.ResourceImporter{
			State: resourceApigeeEnvironmentKeyvaluemapsEntriesImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"env_keyvaluemap_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The Apigee environment keyvalumaps Id associated with the Apigee environment,
in the format 'organizations/{{org_name}}/environments/{{env_name}}/keyvaluemaps/{{keyvaluemap_name}}'.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. Resource URI that can be used to identify the scope of the key value map entries.`,
			},
			"value": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Required. Data or payload that is being retrieved and associated with the unique key.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceApigeeEnvironmentKeyvaluemapsEntriesCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	nameProp, err := expandApigeeEnvironmentKeyvaluemapsEntriesName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !tpgresource.IsEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	valueProp, err := expandApigeeEnvironmentKeyvaluemapsEntriesValue(d.Get("value"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("value"); !tpgresource.IsEmptyValue(reflect.ValueOf(valueProp)) && (ok || !reflect.DeepEqual(v, valueProp)) {
		obj["value"] = valueProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_keyvaluemap_id}}/entries")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new EnvironmentKeyvaluemapsEntries: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating EnvironmentKeyvaluemapsEntries: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "{{env_keyvaluemap_id}}/entries/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating EnvironmentKeyvaluemapsEntries %q: %#v", d.Id(), res)

	return resourceApigeeEnvironmentKeyvaluemapsEntriesRead(d, meta)
}

func resourceApigeeEnvironmentKeyvaluemapsEntriesRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_keyvaluemap_id}}/entries/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("ApigeeEnvironmentKeyvaluemapsEntries %q", d.Id()))
	}

	if err := d.Set("name", flattenApigeeEnvironmentKeyvaluemapsEntriesName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvironmentKeyvaluemapsEntries: %s", err)
	}
	if err := d.Set("value", flattenApigeeEnvironmentKeyvaluemapsEntriesValue(res["value"], d, config)); err != nil {
		return fmt.Errorf("Error reading EnvironmentKeyvaluemapsEntries: %s", err)
	}

	return nil
}

func resourceApigeeEnvironmentKeyvaluemapsEntriesDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{ApigeeBasePath}}{{env_keyvaluemap_id}}/entries/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting EnvironmentKeyvaluemapsEntries %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "EnvironmentKeyvaluemapsEntries")
	}

	log.Printf("[DEBUG] Finished deleting EnvironmentKeyvaluemapsEntries %q: %#v", d.Id(), res)
	return nil
}

func resourceApigeeEnvironmentKeyvaluemapsEntriesImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)

	// current import_formats cannot import fields with forward slashes in their value
	if err := tpgresource.ParseImportId([]string{
		"(?P<env_keyvaluemap_id>.+)/entries/(?P<name>.+)",
		"(?P<env_keyvaluemap_id>.+)/(?P<name>.+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "{{env_keyvaluemap_id}}/entries/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenApigeeEnvironmentKeyvaluemapsEntriesName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenApigeeEnvironmentKeyvaluemapsEntriesValue(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandApigeeEnvironmentKeyvaluemapsEntriesName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandApigeeEnvironmentKeyvaluemapsEntriesValue(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}