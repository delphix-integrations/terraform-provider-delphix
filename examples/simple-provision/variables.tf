variable "dct_hostname" {
  type  = string
  description = "dct hostname config file [default: workspace variable set]"  
}

variable "dct_api_key" {
  type  = string
  description = "dct api key config file [default: workspace variable set]"
}

variable "source_data_id_1" {
  description = "Name or ID of the VDB or Data Source to provision from. [User Defined]"
}