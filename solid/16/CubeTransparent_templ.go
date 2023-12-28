// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CubeTransparentIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M7.6279 1.34882C7.85847 1.21706 8.14153 1.21706 8.3721 1.34882L9.61898 2.06132C9.97862 2.26683 10.1036 2.72497 9.89806 3.0846C9.69255 3.44424 9.23441 3.56919 8.87477 3.36368L8 2.86381L7.12523 3.36368C6.76559 3.56919 6.30745 3.44424 6.10194 3.0846C5.89644 2.72497 6.02138 2.26683 6.38102 2.06132L7.6279 1.34882ZM4.65118 3.91361C4.85669 4.27325 4.73174 4.73139 4.3721 4.9369L4.26167 5L4.3721 5.0631C4.73174 5.26861 4.85669 5.72675 4.65118 6.08639C4.44568 6.44603 3.98753 6.57098 3.6279 6.36547L3.49886 6.29173C3.47721 6.68653 3.15021 7 2.75 7C2.33579 7 2 6.66421 2 6.25V5C2 4.73086 2.14421 4.48235 2.3779 4.34882L3.6279 3.63453C3.98753 3.42902 4.44567 3.55397 4.65118 3.91361ZM11.3488 3.91361C11.5543 3.55397 12.0125 3.42902 12.3721 3.63453L13.6221 4.34882C13.8558 4.48235 14 4.73086 14 5V6.25C14 6.66421 13.6642 7 13.25 7C12.8498 7 12.5228 6.68653 12.5011 6.29173L12.3721 6.36547C12.0125 6.57098 11.5543 6.44603 11.3488 6.08639C11.1433 5.72675 11.2683 5.26861 11.6279 5.0631L11.7383 5L11.6279 4.9369C11.2683 4.73139 11.1433 4.27325 11.3488 3.91361ZM6.10194 6.9154C6.30745 6.55576 6.76559 6.43081 7.12523 6.63632L8 7.13619L8.87477 6.63632C9.23441 6.43081 9.69255 6.55576 9.89806 6.9154C10.1036 7.27503 9.97862 7.73317 9.61898 7.93868L8.75 8.43524V9.25C8.75 9.66421 8.41421 10 8 10C7.58579 10 7.25 9.66421 7.25 9.25V8.43524L6.38102 7.93868C6.02138 7.73317 5.89643 7.27503 6.10194 6.9154ZM2.75 9C3.16421 9 3.5 9.33579 3.5 9.75V10.5648L4.3721 11.0631C4.73174 11.2686 4.85669 11.7268 4.65118 12.0864C4.44568 12.446 3.98753 12.571 3.6279 12.3655L2.3779 11.6512C2.14421 11.5177 2 11.2691 2 11V9.75C2 9.33579 2.33579 9 2.75 9ZM13.25 9C13.6642 9 14 9.33579 14 9.75V11C14 11.2691 13.8558 11.5177 13.6221 11.6512L12.3721 12.3655C12.0125 12.571 11.5543 12.446 11.3488 12.0864C11.1433 11.7268 11.2683 11.2686 11.6279 11.0631L12.5 10.5648V9.75C12.5 9.33579 12.8358 9 13.25 9ZM8.74886 12.7083L8.87477 12.6363C9.23441 12.4308 9.69255 12.5558 9.89806 12.9154C10.1036 13.275 9.97862 13.7332 9.61898 13.9387L8.3721 14.6512C8.14153 14.7829 7.85847 14.7829 7.6279 14.6512L6.38102 13.9387C6.02138 13.7332 5.89643 13.275 6.10194 12.9154C6.30745 12.5558 6.76559 12.4308 7.12523 12.6363L7.25114 12.7083C7.27279 12.3135 7.59979 12 8 12C8.40021 12 8.72721 12.3135 8.74886 12.7083Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
