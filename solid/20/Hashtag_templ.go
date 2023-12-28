// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func HashtagIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M9.49306 2.85248C9.54965 2.44215 9.2629 2.06363 8.85257 2.00703C8.44224 1.95044 8.06372 2.23719 8.00712 2.64752L7.54471 6H4.19831C3.7841 6 3.44831 6.33579 3.44831 6.75C3.44831 7.16421 3.7841 7.5 4.19831 7.5H7.33782L6.64816 12.5H3.30176C2.88754 12.5 2.55176 12.8358 2.55176 13.25C2.55176 13.6642 2.88754 14 3.30176 14H6.44127L6.00713 17.1475C5.95053 17.5578 6.23728 17.9364 6.64761 17.993C7.05794 18.0496 7.43646 17.7628 7.49306 17.3525L7.95547 14H10.9413L10.5071 17.1475C10.4505 17.5578 10.7373 17.9364 11.1476 17.993C11.5579 18.0496 11.9365 17.7628 11.9931 17.3525L12.4555 14H15.8018C16.216 14 16.5518 13.6642 16.5518 13.25C16.5518 12.8358 16.216 12.5 15.8018 12.5H12.6624L13.352 7.5H16.6983C17.1125 7.5 17.4483 7.16421 17.4483 6.75C17.4483 6.33579 17.1125 6 16.6983 6H13.5589L13.9931 2.85248C14.0497 2.44215 13.7629 2.06363 13.3526 2.00703C12.9422 1.95044 12.5637 2.23719 12.5071 2.64752L12.0447 6H9.05892L9.49306 2.85248ZM8.85202 7.5L8.16236 12.5H11.1482L11.8378 7.5H8.85202Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}