package guideme

import (
	"mime/multipart"

	organization "guide.me/gen/organization"
)

// OrganizationMultiAddDecoderFunc implements the multipart decoder for service
// "organization" endpoint "multi_add". The decoder must populate the argument
// p after encoding.
func OrganizationMultiAddDecoderFunc(mr *multipart.Reader, p *[]*organization.Organization) error {
	// Add multipart request decoder logic here
	return nil
}

// OrganizationMultiAddEncoderFunc implements the multipart encoder for service
// "organization" endpoint "multi_add".
func OrganizationMultiAddEncoderFunc(mw *multipart.Writer, p []*organization.Organization) error {
	// Add multipart request encoder logic here
	return nil
}

// OrganizationMultiUpdateDecoderFunc implements the multipart decoder for
// service "organization" endpoint "multi_update". The decoder must populate
// the argument p after encoding.
func OrganizationMultiUpdateDecoderFunc(mr *multipart.Reader, p **organization.MultiUpdatePayload) error {
	// Add multipart request decoder logic here
	return nil
}

// OrganizationMultiUpdateEncoderFunc implements the multipart encoder for
// service "organization" endpoint "multi_update".
func OrganizationMultiUpdateEncoderFunc(mw *multipart.Writer, p *organization.MultiUpdatePayload) error {
	// Add multipart request encoder logic here
	return nil
}
