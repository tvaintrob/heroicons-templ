// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GlobeAmericasIcon(attrs templ.Attributes) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 24 24\" class=\"w-6 h-6\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12 2.25C6.61522 2.25 2.25 6.61522 2.25 12C2.25 17.3848 6.61522 21.75 12 21.75C17.3848 21.75 21.75 17.3848 21.75 12C21.75 6.61522 17.3848 2.25 12 2.25ZM6.26197 6.0723C4.71293 7.57208 3.75 9.67359 3.75 12C3.75 16.5563 7.44365 20.25 12 20.25C16.5563 20.25 20.25 16.5563 20.25 12C20.25 9.24461 18.8992 6.80472 16.8237 5.3064C16.4863 5.84545 16.0374 6.30831 15.5056 6.66289L14.2499 7.5L14.4145 7.82918C14.6835 8.3671 14.2923 9 13.6909 9C13.5653 9 13.4414 8.97076 13.3291 8.91459L12.7252 8.61262C12.2921 8.39607 11.769 8.48095 11.4266 8.82336L11.2954 8.9545C10.8561 9.39384 10.8561 10.1062 11.2954 10.5455L11.5905 10.8406C11.8474 11.0975 12.2126 11.2146 12.571 11.1548L13.7411 10.9598C14.0641 10.906 14.3946 10.9956 14.6462 11.2053L15.9755 12.313C16.2962 12.5802 16.4356 13.0073 16.3344 13.4122C15.9519 14.9419 15.1609 16.339 14.046 17.4539L13.3233 18.1766C12.9809 18.519 12.4578 18.6039 12.0247 18.3874L11.8718 18.3109C11.4907 18.1204 11.2499 17.7308 11.2499 17.3047V16.216C11.2499 15.9176 11.1314 15.6315 10.9204 15.4205L9.57328 14.0734C9.23087 13.731 9.14599 13.2079 9.36254 12.7747L9.74992 12L8.10954 10.3596C7.22527 9.47535 6.6394 8.33689 6.43381 7.10337L6.26197 6.0723Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
