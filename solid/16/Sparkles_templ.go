// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SparklesIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5 4C5.36246 4 5.67306 4.25922 5.7379 4.61584L5.99021 6.00355C6.08335 6.51578 6.48422 6.91665 6.99645 7.00979L8.38416 7.2621C8.74078 7.32694 9 7.63754 9 8C9 8.36246 8.74078 8.67306 8.38416 8.7379L6.99645 8.99021C6.48422 9.08335 6.08335 9.48422 5.99021 9.99645L5.7379 11.3842C5.67306 11.7408 5.36246 12 5 12C4.63754 12 4.32694 11.7408 4.2621 11.3842L4.00979 9.99645C3.91665 9.48422 3.51578 9.08335 3.00355 8.99021L1.61584 8.7379C1.25922 8.67306 1 8.36246 1 8C1 7.63754 1.25922 7.32694 1.61584 7.2621L3.00355 7.00979C3.51578 6.91665 3.91665 6.51578 4.00979 6.00355L4.2621 4.61584C4.32694 4.25922 4.63754 4 5 4Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12 1C12.3349 1 12.6291 1.22198 12.7211 1.54396L12.9159 2.2256C13.0345 2.64086 13.3591 2.96546 13.7744 3.0841L14.456 3.27886C14.778 3.37085 15 3.66514 15 4C15 4.33486 14.778 4.62915 14.456 4.72114L13.7744 4.9159C13.3591 5.03454 13.0345 5.35914 12.9159 5.7744L12.7211 6.45604C12.6291 6.77802 12.3349 7 12 7C11.6651 7 11.3709 6.77802 11.2789 6.45604L11.0841 5.7744C10.9655 5.35914 10.6409 5.03454 10.2256 4.9159L9.54396 4.72114C9.22198 4.62915 9 4.33486 9 4C9 3.66514 9.22198 3.37085 9.54396 3.27886L10.2256 3.0841C10.6409 2.96546 10.9655 2.64086 11.0841 2.2256L11.2789 1.54396C11.3709 1.22198 11.6651 1 12 1Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M10 11C10.3442 11 10.6441 11.2342 10.7276 11.5681C10.8143 11.9149 11.0851 12.1857 11.4319 12.2724C11.7658 12.3559 12 12.6558 12 13C12 13.3442 11.7658 13.6441 11.4319 13.7276C11.0851 13.8143 10.8143 14.0851 10.7276 14.4319C10.6441 14.7658 10.3442 15 10 15C9.65585 15 9.35586 14.7658 9.27239 14.4319C9.18569 14.0851 8.9149 13.8143 8.5681 13.7276C8.23422 13.6441 8 13.3442 8 13C8 12.6558 8.23422 12.3559 8.5681 12.2724C8.9149 12.1857 9.18569 11.9149 9.27239 11.5681C9.35586 11.2342 9.65585 11 10 11Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
