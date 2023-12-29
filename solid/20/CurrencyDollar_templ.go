// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyDollarIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path d=\"M10.7499 10.8176V13.4324C11.1816 13.3527 11.5745 13.2046 11.8876 12.9999C12.3698 12.6846 12.4999 12.352 12.4999 12.125C12.4999 11.898 12.3698 11.5654 11.8876 11.2501C11.5745 11.0454 11.1816 10.8973 10.7499 10.8176Z\" fill=\"#0F172A\"></path> <path d=\"M8.32961 8.61947C8.38337 8.67543 8.44464 8.73053 8.51404 8.78416C8.72197 8.94484 8.97355 9.06777 9.25 9.1469V6.60315C9.17545 6.62449 9.10271 6.64901 9.03215 6.6766C8.98721 6.69417 8.94315 6.71299 8.90007 6.73302C8.75996 6.79816 8.63019 6.87614 8.51404 6.96589C8.13658 7.25757 8 7.59253 8 7.87503C8 8.05887 8.05784 8.26493 8.20228 8.46683C8.23898 8.51812 8.28126 8.56915 8.32961 8.61947Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M18 10C18 14.4183 14.4183 18 10 18C5.58172 18 2 14.4183 2 10C2 5.58172 5.58172 2 10 2C14.4183 2 18 5.58172 18 10ZM9.99994 4C10.4142 4 10.7499 4.33579 10.7499 4.75V5.06584C11.3423 5.17106 11.9182 5.40427 12.4031 5.77893C12.8293 6.10829 13.1467 6.51836 13.3282 6.97896C13.4801 7.36432 13.2908 7.79985 12.9055 7.95174C12.5201 8.10363 12.0846 7.91437 11.9327 7.52901C11.8599 7.34437 11.72 7.14675 11.4859 6.96586C11.278 6.80519 11.0264 6.68225 10.7499 6.60312V9.29944C11.448 9.39233 12.1327 9.61819 12.7085 9.99467C13.4955 10.5093 13.9999 11.2644 13.9999 12.125C13.9999 12.9856 13.4955 13.7407 12.7085 14.2553C12.1327 14.6318 11.448 14.8577 10.7499 14.9506V15.25C10.7499 15.6642 10.4142 16 9.99994 16C9.58573 16 9.24994 15.6642 9.24994 15.25V14.9506C8.55186 14.8577 7.8672 14.6318 7.29141 14.2553C6.80887 13.9398 6.4337 13.5376 6.21337 13.0672C6.0377 12.692 6.19937 12.2455 6.57449 12.0699C6.9496 11.8942 7.39611 12.0559 7.57178 12.431C7.65258 12.6035 7.81692 12.8067 8.11229 12.9999C8.42537 13.2046 8.8183 13.3526 9.24994 13.4324V10.6842C8.65762 10.5789 8.08167 10.3457 7.59681 9.97107C6.90033 9.43288 6.49994 8.68017 6.49994 7.875C6.49994 7.06983 6.90034 6.31712 7.59681 5.77893C8.08167 5.40427 8.65762 5.17106 9.24994 5.06584V4.75C9.24994 4.33579 9.58573 4 9.99994 4Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
