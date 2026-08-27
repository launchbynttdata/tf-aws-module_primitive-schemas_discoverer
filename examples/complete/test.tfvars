logical_product_family  = "lp"
logical_product_service = "lps"
class_env               = "dev"
instance_env            = 0
instance_resource       = 0

resource_names_map = {
  eventbus = {
    name       = "eventbus"
    max_length = 64
  }
}

description = "Complete example schema discoverer."

tags = {
  Environment = "test"
  Terraform   = "true"
}
