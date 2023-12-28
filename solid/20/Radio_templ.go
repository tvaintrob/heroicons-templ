// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M17.4497 3.47294C17.849 3.36279 18.0834 2.9498 17.9732 2.5505C17.8631 2.1512 17.4501 1.9168 17.0508 2.02695L5.31247 5.2651C4.47281 5.36069 3.64077 5.48165 2.81717 5.62721C1.74546 5.81661 1 6.75574 1 7.81601V15.75C1 16.9926 2.00736 18 3.25 18H16.75C17.9926 18 19 16.9926 19 15.75V7.81601C19 6.75573 18.2545 5.81661 17.1828 5.62721C15.417 5.31513 13.6124 5.11612 11.777 5.03783L17.4497 3.47294ZM16 9.5C16 10.3284 15.3284 11 14.5 11C13.6716 11 13 10.3284 13 9.5C13 8.67157 13.6716 8 14.5 8C15.3284 8 16 8.67157 16 9.5ZM14.5 16C15.3284 16 16 15.3284 16 14.5C16 13.6716 15.3284 13 14.5 13C13.6716 13 13 13.6716 13 14.5C13 15.3284 13.6716 16 14.5 16ZM5.24023 11C5.24023 10.5858 5.57602 10.25 5.99023 10.25H6.00023C6.41445 10.25 6.75023 10.5858 6.75023 11V11.01C6.75023 11.4242 6.41445 11.76 6.00023 11.76H5.99023C5.57602 11.76 5.24023 11.4242 5.24023 11.01V11ZM7.99023 10.25C7.57602 10.25 7.24023 10.5858 7.24023 11V11.01C7.24023 11.4242 7.57602 11.76 7.99023 11.76H8.00023C8.41445 11.76 8.75023 11.4242 8.75023 11.01V11C8.75023 10.5858 8.41445 10.25 8.00023 10.25H7.99023ZM6.24023 8.75C6.24023 8.33579 6.57602 8 6.99023 8H7.00023C7.41445 8 7.75023 8.33579 7.75023 8.75V8.76C7.75023 9.17421 7.41445 9.51 7.00023 9.51H6.99023C6.57602 9.51 6.24023 9.17421 6.24023 8.76V8.75ZM9.8233 9.16949C9.53041 8.8766 9.05553 8.8766 8.76264 9.16949L8.75557 9.17656C8.46267 9.46945 8.46267 9.94433 8.75557 10.2372L8.76264 10.2443C9.05553 10.5372 9.53041 10.5372 9.8233 10.2443L9.83037 10.2372C10.1233 9.94433 10.1233 9.46945 9.83037 9.17656L9.8233 9.16949ZM10.25 11.25C10.6642 11.25 11 11.5858 11 12V12.01C11 12.4242 10.6642 12.76 10.25 12.76H10.24C9.82579 12.76 9.49 12.4242 9.49 12.01V12C9.49 11.5858 9.82579 11.25 10.24 11.25H10.25ZM9.83051 14.8334C10.1234 14.5405 10.1234 14.0657 9.83051 13.7728L9.82344 13.7657C9.53055 13.4728 9.05567 13.4728 8.76278 13.7657L8.75571 13.7728C8.46281 14.0657 8.46281 14.5405 8.75571 14.8334L8.76278 14.8405C9.05567 15.1334 9.53055 15.1334 9.82344 14.8405L9.83051 14.8334ZM6.24023 15.25C6.24023 14.8358 6.57602 14.5 6.99023 14.5H7.00023C7.41445 14.5 7.75023 14.8358 7.75023 15.25V15.26C7.75023 15.6742 7.41445 16.01 7.00023 16.01H6.99023C6.57602 16.01 6.24023 15.6742 6.24023 15.26V15.25ZM5.22711 13.7657C4.93421 13.4728 4.45934 13.4728 4.16645 13.7657L4.15938 13.7728C3.86648 14.0656 3.86648 14.5405 4.15938 14.8334L4.16645 14.8405C4.45934 15.1334 4.93421 15.1334 5.22711 14.8405L5.23418 14.8334C5.52707 14.5405 5.52707 14.0656 5.23418 13.7728L5.22711 13.7657ZM3.75 11.25C4.16421 11.25 4.5 11.5858 4.5 12V12.01C4.5 12.4242 4.16421 12.76 3.75 12.76H3.74C3.32579 12.76 2.99 12.4242 2.99 12.01V12C2.99 11.5858 3.32579 11.25 3.74 11.25H3.75ZM5.23432 10.2372C5.52721 9.94435 5.52721 9.46947 5.23432 9.17658L5.22725 9.16951C4.93435 8.87661 4.45948 8.87661 4.16659 9.16951L4.15952 9.17658C3.86662 9.46947 3.86662 9.94435 4.15952 10.2372L4.16659 10.2443C4.45948 10.5372 4.93435 10.5372 5.22725 10.2443L5.23432 10.2372ZM7.24023 13C7.24023 12.5858 7.57602 12.25 7.99023 12.25H8.00023C8.41445 12.25 8.75023 12.5858 8.75023 13V13.01C8.75023 13.4242 8.41445 13.76 8.00023 13.76H7.99023C7.57602 13.76 7.24023 13.4242 7.24023 13.01V13ZM5.99023 12.25C5.57602 12.25 5.24023 12.5858 5.24023 13V13.01C5.24023 13.4242 5.57602 13.76 5.99023 13.76H6.00023C6.41445 13.76 6.75023 13.4242 6.75023 13.01V13C6.75023 12.5858 6.41445 12.25 6.00023 12.25H5.99023Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
