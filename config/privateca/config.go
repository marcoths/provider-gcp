package privateca

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom
// ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_privateca_ca_pool_iam_member", func(r *config.Resource) {
		r.References["ca_pool"] = config.Reference{
			Type:      "CAPool",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_privateca_certificate_template_iam_member", func(r *config.Resource) {
		r.References["certificate_template"] = config.Reference{
			Type:      "CertificateTemplate",
			Extractor: common.ExtractResourceIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_privateca_certificate_authority", func(r *config.Resource) {
		r.References["pool"] = config.Reference{
			Type: "CAPool",
		}
	})

	p.AddResourceConfigurator("google_privateca_certificate", func(r *config.Resource) {
		r.References["pool"] = config.Reference{
			Type: "CAPool",
		}
		r.TerraformResource.Schema["config"].Elem.(*schema.Resource).
			Schema["public_key"].Elem.(*schema.Resource).
			Schema["key"].Sensitive = true
	})
}
