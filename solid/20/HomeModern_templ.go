// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func HomeModernIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M14.9156 2.40399C15.1067 2.77148 14.9637 3.22432 14.5962 3.41541L14 3.72542V17C14 17.5523 13.5523 18 13 18H10.74C10.3258 18 9.99 17.6642 9.99 17.25V13.75C9.99 13.3358 9.65421 13 9.24 13H6.75C6.33579 13 6 13.3358 6 13.75V17.25C6 17.6642 5.66421 18 5.25 18H1.75C1.33579 18 1 17.6642 1 17.25C1 16.8358 1.33579 16.5 1.75 16.5H2V9.95726C1.65296 10.0803 1.25946 9.93201 1.08474 9.59602C0.893644 9.22852 1.03664 8.77569 1.40414 8.58459L2 8.27474V5.75C2 5.33579 2.33579 5 2.75 5C3.16421 5 3.5 5.33579 3.5 5.75V7.49474L13.9041 2.08459C14.2716 1.89349 14.7245 2.03649 14.9156 2.40399Z\" fill=\"#0F172A\"></path> <path d=\"M15.8615 8.56923C16.0854 8.43325 16.3642 8.42404 16.5967 8.54493L18.596 9.58459C18.9635 9.77569 19.1065 10.2285 18.9154 10.596C18.7407 10.932 18.3473 11.0803 18.0003 10.9573V16.5H18.25C18.6642 16.5 19 16.8358 19 17.25C19 17.6642 18.6642 18 18.25 18H16.25C16.0511 18 15.8603 17.921 15.7196 17.7803C15.579 17.6396 15.5 17.4489 15.5 17.2499L15.5006 9.21028C15.5007 8.94824 15.6375 8.70521 15.8615 8.56923Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
