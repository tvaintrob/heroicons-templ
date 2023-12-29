// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CalculatorIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M6.32022 1.82741C8.18374 1.61114 10.079 1.5 12 1.5C13.921 1.5 15.8163 1.61114 17.6798 1.82741C19.1772 2.00119 20.25 3.28722 20.25 4.75699V19.5C20.25 21.1569 18.9069 22.5 17.25 22.5H6.75C5.09315 22.5 3.75 21.1569 3.75 19.5V4.75699C3.75 3.28722 4.82283 2.00119 6.32022 1.82741ZM7.5 11.25C7.5 10.8358 7.83579 10.5 8.25 10.5H8.2575C8.67171 10.5 9.0075 10.8358 9.0075 11.25V11.2575C9.0075 11.6717 8.67171 12.0075 8.2575 12.0075H8.25C7.83579 12.0075 7.5 11.6717 7.5 11.2575V11.25ZM8.25 12.75C7.83579 12.75 7.5 13.0858 7.5 13.5V13.5075C7.5 13.9217 7.83579 14.2575 8.25 14.2575H8.2575C8.67171 14.2575 9.0075 13.9217 9.0075 13.5075V13.5C9.0075 13.0858 8.67171 12.75 8.2575 12.75H8.25ZM7.5 15.75C7.5 15.3358 7.83579 15 8.25 15H8.2575C8.67171 15 9.0075 15.3358 9.0075 15.75V15.7575C9.0075 16.1717 8.67171 16.5075 8.2575 16.5075H8.25C7.83579 16.5075 7.5 16.1717 7.5 15.7575V15.75ZM8.25 17.25C7.83579 17.25 7.5 17.5858 7.5 18V18.0075C7.5 18.4217 7.83579 18.7575 8.25 18.7575H8.2575C8.67171 18.7575 9.0075 18.4217 9.0075 18.0075V18C9.0075 17.5858 8.67171 17.25 8.2575 17.25H8.25ZM9.99756 11.25C9.99756 10.8358 10.3333 10.5 10.7476 10.5H10.7551C11.1693 10.5 11.5051 10.8358 11.5051 11.25V11.2575C11.5051 11.6717 11.1693 12.0075 10.7551 12.0075H10.7476C10.3333 12.0075 9.99756 11.6717 9.99756 11.2575V11.25ZM10.7476 12.75C10.3333 12.75 9.99756 13.0858 9.99756 13.5V13.5075C9.99756 13.9217 10.3333 14.2575 10.7476 14.2575H10.7551C11.1693 14.2575 11.5051 13.9217 11.5051 13.5075V13.5C11.5051 13.0858 11.1693 12.75 10.7551 12.75H10.7476ZM9.99756 15.75C9.99756 15.3358 10.3333 15 10.7476 15H10.7551C11.1693 15 11.5051 15.3358 11.5051 15.75V15.7575C11.5051 16.1717 11.1693 16.5075 10.7551 16.5075H10.7476C10.3333 16.5075 9.99756 16.1717 9.99756 15.7575V15.75ZM10.7476 17.25C10.3333 17.25 9.99756 17.5858 9.99756 18V18.0075C9.99756 18.4217 10.3333 18.7575 10.7476 18.7575H10.7551C11.1693 18.7575 11.5051 18.4217 11.5051 18.0075V18C11.5051 17.5858 11.1693 17.25 10.7551 17.25H10.7476ZM12.5024 11.25C12.5024 10.8358 12.8382 10.5 13.2524 10.5H13.2599C13.6742 10.5 14.0099 10.8358 14.0099 11.25V11.2575C14.0099 11.6717 13.6742 12.0075 13.2599 12.0075H13.2524C12.8382 12.0075 12.5024 11.6717 12.5024 11.2575V11.25ZM13.2524 12.75C12.8382 12.75 12.5024 13.0858 12.5024 13.5V13.5075C12.5024 13.9217 12.8382 14.2575 13.2524 14.2575H13.2599C13.6742 14.2575 14.0099 13.9217 14.0099 13.5075V13.5C14.0099 13.0858 13.6742 12.75 13.2599 12.75H13.2524ZM12.5024 15.75C12.5024 15.3358 12.8382 15 13.2524 15H13.2599C13.6742 15 14.0099 15.3358 14.0099 15.75V15.7575C14.0099 16.1717 13.6742 16.5075 13.2599 16.5075H13.2524C12.8382 16.5075 12.5024 16.1717 12.5024 15.7575V15.75ZM13.2524 17.25C12.8382 17.25 12.5024 17.5858 12.5024 18V18.0075C12.5024 18.4217 12.8382 18.7575 13.2524 18.7575H13.2599C13.6742 18.7575 14.0099 18.4217 14.0099 18.0075V18C14.0099 17.5858 13.6742 17.25 13.2599 17.25H13.2524ZM15 11.25C15 10.8358 15.3358 10.5 15.75 10.5H15.7575C16.1717 10.5 16.5075 10.8358 16.5075 11.25V11.2575C16.5075 11.6717 16.1717 12.0075 15.7575 12.0075H15.75C15.3358 12.0075 15 11.6717 15 11.2575V11.25ZM15.75 12.75C15.3358 12.75 15 13.0858 15 13.5V13.5075C15 13.9217 15.3358 14.2575 15.75 14.2575H15.7575C16.1717 14.2575 16.5075 13.9217 16.5075 13.5075V13.5C16.5075 13.0858 16.1717 12.75 15.7575 12.75H15.75ZM7.5 6.75C7.5 6.33579 7.83579 6 8.25 6H15.75C16.1642 6 16.5 6.33579 16.5 6.75V7.5C16.5 7.91421 16.1642 8.25 15.75 8.25H8.25C7.83579 8.25 7.5 7.91421 7.5 7.5V6.75ZM16.5 15.75C16.5 15.3358 16.1642 15 15.75 15C15.3358 15 15 15.3358 15 15.75V18C15 18.4142 15.3358 18.75 15.75 18.75C16.1642 18.75 16.5 18.4142 16.5 18V15.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
