// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"strconv"
)

func NotesForm(count int) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"max-w-3xl mx-auto drop-shadow-2xl relative\"><span class=\"absolute top-2 right-2 text-xl\">Notes Count: <span class=\"text-large\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(strconv.Itoa(count))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/note_form.templ`, Line: 6, Col: 115}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></span><h2 class=\"text-xl mx-auto text-center\">New Note</h2><form class=\"flex flex-col gap-4\" hx-post=\"/notes\" hx-swap=\"innerHTML\" hx-target=\"#main\" hx-indicator=\"#loading\"><div class=\"sm:col-span-4\"><label for=\"title\" class=\"block text-sm font-medium leading-6 text-gray-900\">Title</label><div class=\"mt-2\"><input type=\"text\" name=\"title\" id=\"title\" autocomplete=\"title\" class=\"block flex-1 border-2 rounded bg-transparent py-1.5 pl-1 text-gray-900 placeholder:text-gray-400 focus:ring-0 sm:text-sm sm:leading-6\" placeholder=\"My note title\"></div></div><div class=\"col-span-full\"><label for=\"content\" class=\"block text-sm font-medium leading-6 text-gray-900\">Content</label><div class=\"mt-2\"><textarea id=\"content\" name=\"body\" rows=\"3\" class=\"block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6\"></textarea></div></div><div class=\"mt-6 flex items-center justify-end gap-x-6\"><button type=\"submit\" class=\"rounded-md bg-indigo-600 px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600\">Save</button></div></form><div class=\"text-2xl htmx-indicator\" id=\"loading\">Creating Note 🕐 </div></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
