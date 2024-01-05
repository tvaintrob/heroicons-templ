// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ScaleIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 16 16\" class=\"w-4 h-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8.75 2.5C8.75 2.08579 8.41421 1.75 8 1.75C7.58579 1.75 7.25 2.08579 7.25 2.5V3.00849C5.67898 3.04408 4.13508 3.1912 2.62601 3.44235C2.21742 3.51035 1.94132 3.89671 2.00932 4.3053C2.07732 4.71389 2.46367 4.99 2.87227 4.922L3.00168 4.90074L1.81432 9.6502C1.73555 9.96527 1.86852 10.2952 2.14373 10.4677C2.68226 10.8051 3.31949 11 3.99995 11C4.68041 11 5.31763 10.8051 5.85616 10.4677C6.13138 10.2952 6.26435 9.96527 6.18558 9.6502L4.93541 4.64952C5.69905 4.5746 6.47091 4.52738 7.25 4.50889V12L7.25009 12.0117C6.36827 12.0394 5.49846 12.1157 4.64327 12.238C4.23323 12.2967 3.94838 12.6766 4.00704 13.0867C4.06571 13.4967 4.44566 13.7816 4.8557 13.7229C5.88227 13.576 6.93206 13.5 8.00004 13.5C9.06803 13.5 10.1178 13.576 11.1444 13.7229C11.5544 13.7816 11.9344 13.4967 11.993 13.0867C12.0517 12.6766 11.7669 12.2967 11.3568 12.238C10.5016 12.1157 9.63176 12.0394 8.74991 12.0117L8.75 12V4.50889C9.52877 4.52737 10.3003 4.57456 11.0637 4.64942L9.81346 9.6502C9.73469 9.96527 9.86766 10.2952 10.1429 10.4677C10.6814 10.8051 11.3186 11 11.9991 11C12.6796 11 13.3168 10.8051 13.8553 10.4677C14.1305 10.2952 14.2635 9.96527 14.1847 9.6502L12.9973 4.90057C13.0408 4.90762 13.0843 4.91477 13.1278 4.922C13.5364 4.99 13.9227 4.71389 13.9907 4.3053C14.0587 3.89671 13.7826 3.51035 13.374 3.44235C11.865 3.1912 10.321 3.04408 8.75 3.00849V2.5ZM3.41937 9.41464C3.60278 9.47014 3.79757 9.5 3.99995 9.5C4.20233 9.5 4.39712 9.47014 4.58052 9.41464L3.99995 7.09233L3.41937 9.41464ZM11.9991 9.5C11.7967 9.5 11.6019 9.47014 11.4185 9.41464L11.9991 7.09233L12.5797 9.41464C12.3963 9.47014 12.2015 9.5 11.9991 9.5Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
