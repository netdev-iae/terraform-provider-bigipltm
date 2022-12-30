package bigipcustum

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	bigip "github.com/idrissae/go-bigip-client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Domain name/IP of the BigIP",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_HOST", nil),
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Management Port to connect to Bigip",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_PORT", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Username with API access to the BigIP",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_USER", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The user's password. Leave empty if using token_value",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_PASSWORD", nil),
			},
			"token_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A token generated outside the provider, in place of password",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_TOKEN_VALUE", nil),
			},
			"token_auth": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable to use an external authentication source (LDAP, TACACS, etc)",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_TOKEN_AUTH", true),
			},
			"validate_certs_disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set to true, Disables TLS certificate check on BIG-IP. Default : True",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_VERIFY_CERT_DISABLE", true),
			},
			"trusted_cert_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Valid Trusted Certificate path",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_TRUSTED_CERT_PATH", nil),
			},
			"teem_disable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If this flag set to true,sending telemetry data to TEEM will be disabled",
				DefaultFunc: schema.EnvDefaultFunc("TEEM_DISABLE", false),
			},
			"login_ref": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Login reference for token authentication (see BIG-IP REST docs for details)",
				DefaultFunc: schema.EnvDefaultFunc("BIGIP_LOGIN_REF", "tmos"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"bigipcustum_monitor_custum": resourceBigipLtmMonitor(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"bigipcustum_monitor_custum": dataSourceBigipLtmMonitor(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	config := &bigip.Config{
		Address:           d.Get("address").(string),
		Port:              d.Get("port").(string),
		Username:          d.Get("username").(string),
		Password:          d.Get("password").(string),
		Token:             d.Get("token_value").(string),
		CertVerifyDisable: d.Get("validate_certs_disable").(bool),
	}
	if d.Get("token_auth").(bool) {
		config.LoginReference = d.Get("login_ref").(string)
	}
	if !d.Get("validate_certs_disable").(bool) {
		if d.Get("trusted_cert_path").(string) == "" {
			return nil, diag.FromErr(fmt.Errorf("Valid Trust Certificate path not provided using :%+v ", "trusted_cert_path"))
		}
		config.TrustedCertificate = d.Get("trusted_cert_path").(string)
	}
	cfg, err := Client(config)
	if err != nil {
		return cfg, diag.FromErr(err)
	}
	if cfg != nil {
		//cfg.UserAgent = fmt.Sprintf("Terraform/%s", terraformVersion)
		//cfg.UserAgent += fmt.Sprintf("/terraform-provider-bigip/%s", getVersion())
		cfg.Teem = d.Get("teem_disable").(bool)
		cfg.Transport.TLSClientConfig.InsecureSkipVerify = d.Get("validate_certs_disable").(bool)
	}
	return cfg, diag.FromErr(err)
}

// Convert schema.Set to a slice of strings
func setToStringSlice(s *schema.Set) []string {
	list := make([]string, s.Len())
	for i, v := range s.List() {
		list[i] = v.(string)
	}
	return list
}
