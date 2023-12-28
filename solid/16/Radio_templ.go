// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func RadioIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.4371 1.44939C11.6031 1.82888 11.43 2.2711 11.0506 2.43712L7.47827 4H13C14.1046 4 15 4.89543 15 6V12C15 13.1046 14.1046 14 13 14H3C1.89543 14 1 13.1046 1 12V6C1 4.89543 1.89543 4 3 4H3.73593L10.4493 1.06289C10.8288 0.896864 11.271 1.06991 11.4371 1.44939ZM12 8C12.5523 8 13 7.55228 13 7C13 6.44772 12.5523 6 12 6C11.4477 6 11 6.44772 11 7C11 7.55228 11.4477 8 12 8ZM6.75 6.75C6.75 7.16421 6.41421 7.5 6 7.5C5.58579 7.5 5.25 7.16421 5.25 6.75C5.25 6.33579 5.58579 6 6 6C6.41421 6 6.75 6.33579 6.75 6.75ZM6 9.75C6.41421 9.75 6.75 9.41421 6.75 9C6.75 8.58579 6.41421 8.25 6 8.25C5.58579 8.25 5.25 8.58579 5.25 9C5.25 9.41421 5.58579 9.75 6 9.75ZM8.32314 8.52462C7.96442 8.73173 7.50572 8.60882 7.29862 8.2501C7.09151 7.89139 7.21442 7.43269 7.57314 7.22559C7.93186 7.01848 8.39055 7.14139 8.59766 7.5001C8.80476 7.85882 8.68186 8.31752 8.32314 8.52462ZM7.29862 9.7501C7.09151 10.1088 7.21442 10.5675 7.57314 10.7746C7.93186 10.9817 8.39055 10.8588 8.59766 10.5001C8.80476 10.1414 8.68186 9.68269 8.32314 9.47559C7.96442 9.26848 7.50572 9.39139 7.29862 9.7501ZM6.75 11.25C6.75 11.6642 6.41421 12 6 12C5.58579 12 5.25 11.6642 5.25 11.25C5.25 10.8358 5.58579 10.5 6 10.5C6.41421 10.5 6.75 10.8358 6.75 11.25ZM3.40213 10.5001C3.60924 10.8588 4.06793 10.9817 4.42665 10.7746C4.78537 10.5675 4.90828 10.1088 4.70117 9.7501C4.49407 9.39139 4.03537 9.26848 3.67665 9.47559C3.31793 9.68269 3.19503 10.1414 3.40213 10.5001ZM3.67665 8.52462C3.31793 8.31752 3.19503 7.85882 3.40213 7.5001C3.60924 7.14139 4.06793 7.01848 4.42665 7.22559C4.78537 7.43269 4.90828 7.89139 4.70117 8.2501C4.49407 8.60882 4.03537 8.73173 3.67665 8.52462ZM12 12C12.5523 12 13 11.5523 13 11C13 10.4477 12.5523 10 12 10C11.4477 10 11 10.4477 11 11C11 11.5523 11.4477 12 12 12Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}