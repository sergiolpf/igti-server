// Code generated by goa v3.1.1, DO NOT EDIT.
//
// organization HTTP server encoders and decoders
//
// Command:
// $ goa gen guide.me/design

package server

import (
	"context"
	"io"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	organization "guide.me/gen/organization"
	organizationviews "guide.me/gen/organization/views"
)

// EncodeListResponse returns an encoder for responses returned by the
// organization list endpoint.
func EncodeListResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(organizationviews.StoredOrganizationCollection)
		enc := encoder(ctx, w)
		body := NewStoredOrganizationResponseTinyCollection(res.Projected)
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// EncodeShowResponse returns an encoder for responses returned by the
// organization show endpoint.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*organizationviews.StoredOrganization)
		w.Header().Set("goa-view", res.View)
		enc := encoder(ctx, w)
		var body interface{}
		switch res.View {
		case "default", "":
			body = NewShowResponseBody(res.Projected)
		case "tiny":
			body = NewShowResponseBodyTiny(res.Projected)
		}
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeShowRequest returns a decoder for requests sent to the organization
// show endpoint.
func DecodeShowRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id   string
			view *string
			err  error

			params = mux.Vars(r)
		)
		id = params["id"]
		viewRaw := r.URL.Query().Get("view")
		if viewRaw != "" {
			view = &viewRaw
		}
		if view != nil {
			if !(*view == "default" || *view == "tiny") {
				err = goa.MergeErrors(err, goa.InvalidEnumValueError("view", *view, []interface{}{"default", "tiny"}))
			}
		}
		if err != nil {
			return nil, err
		}
		payload := NewShowPayload(id, view)

		return payload, nil
	}
}

// EncodeShowError returns an encoder for errors returned by the show
// organization endpoint.
func EncodeShowError(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder, formatter func(err error) goahttp.Statuser) func(context.Context, http.ResponseWriter, error) error {
	encodeError := goahttp.ErrorEncoder(encoder, formatter)
	return func(ctx context.Context, w http.ResponseWriter, v error) error {
		en, ok := v.(ErrorNamer)
		if !ok {
			return encodeError(ctx, w, v)
		}
		switch en.ErrorName() {
		case "not_found":
			res := v.(*organization.OrgNotFound)
			enc := encoder(ctx, w)
			var body interface{}
			if formatter != nil {
				body = formatter(res)
			} else {
				body = NewShowNotFoundResponseBody(res)
			}
			w.Header().Set("goa-error", "not_found")
			w.WriteHeader(http.StatusNotFound)
			return enc.Encode(body)
		default:
			return encodeError(ctx, w, v)
		}
	}
}

// EncodeAddResponse returns an encoder for responses returned by the
// organization add endpoint.
func EncodeAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeAddRequest returns a decoder for requests sent to the organization add
// endpoint.
func DecodeAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body AddRequestBody
			err  error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		err = ValidateAddRequestBody(&body)
		if err != nil {
			return nil, err
		}
		payload := NewAddOrganization(&body)

		return payload, nil
	}
}

// EncodeRemoveResponse returns an encoder for responses returned by the
// organization remove endpoint.
func EncodeRemoveResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeRemoveRequest returns a decoder for requests sent to the organization
// remove endpoint.
func DecodeRemoveRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id string

			params = mux.Vars(r)
		)
		id = params["id"]
		payload := NewRemovePayload(id)

		return payload, nil
	}
}

// EncodeMultiAddResponse returns an encoder for responses returned by the
// organization multi_add endpoint.
func EncodeMultiAddResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.([]string)
		enc := encoder(ctx, w)
		body := res
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeMultiAddRequest returns a decoder for requests sent to the
// organization multi_add endpoint.
func DecodeMultiAddRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload []*organization.Organization
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewOrganizationMultiAddDecoder returns a decoder to decode the multipart
// request for the "organization" service "multi_add" endpoint.
func NewOrganizationMultiAddDecoder(mux goahttp.Muxer, organizationMultiAddDecoderFn OrganizationMultiAddDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v interface{}) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(*[]*organization.Organization)
			if err := organizationMultiAddDecoderFn(mr, p); err != nil {
				return err
			}
			return nil
		})
	}
}

// EncodeMultiUpdateResponse returns an encoder for responses returned by the
// organization multi_update endpoint.
func EncodeMultiUpdateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// DecodeMultiUpdateRequest returns a decoder for requests sent to the
// organization multi_update endpoint.
func DecodeMultiUpdateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var payload *organization.MultiUpdatePayload
		if err := decoder(r).Decode(&payload); err != nil {
			return nil, goa.DecodePayloadError(err.Error())
		}

		return payload, nil
	}
}

// NewOrganizationMultiUpdateDecoder returns a decoder to decode the multipart
// request for the "organization" service "multi_update" endpoint.
func NewOrganizationMultiUpdateDecoder(mux goahttp.Muxer, organizationMultiUpdateDecoderFn OrganizationMultiUpdateDecoderFunc) func(r *http.Request) goahttp.Decoder {
	return func(r *http.Request) goahttp.Decoder {
		return goahttp.EncodingFunc(func(v interface{}) error {
			mr, merr := r.MultipartReader()
			if merr != nil {
				return merr
			}
			p := v.(**organization.MultiUpdatePayload)
			if err := organizationMultiUpdateDecoderFn(mr, p); err != nil {
				return err
			}

			var (
				ids []string
				err error
			)
			ids = r.URL.Query()["ids"]
			if ids == nil {
				err = goa.MergeErrors(err, goa.MissingFieldError("ids", "query string"))
			}
			if err != nil {
				return err
			}
			(*p).Ids = ids
			return nil
		})
	}
}

// marshalOrganizationviewsStoredOrganizationViewToStoredOrganizationResponseTiny
// builds a value of type *StoredOrganizationResponseTiny from a value of type
// *organizationviews.StoredOrganizationView.
func marshalOrganizationviewsStoredOrganizationViewToStoredOrganizationResponseTiny(v *organizationviews.StoredOrganizationView) *StoredOrganizationResponseTiny {
	res := &StoredOrganizationResponseTiny{
		ID:   *v.ID,
		Name: *v.Name,
	}

	return res
}

// unmarshalOrganizationRequestBodyToOrganizationOrganization builds a value of
// type *organization.Organization from a value of type
// *OrganizationRequestBody.
func unmarshalOrganizationRequestBodyToOrganizationOrganization(v *OrganizationRequestBody) *organization.Organization {
	res := &organization.Organization{
		Name: *v.Name,
		URL:  *v.URL,
	}

	return res
}
