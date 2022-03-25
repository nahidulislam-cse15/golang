
DATE:25 MARCH 20222
# GO TEMPLATE
a present format for documet or file

Go templates are a powerful method to customize output however you want, whether you’re creating a web page, sending an e-mail, working with Buffalo, Go-Hugo, or just using some CLI such as kubectl.
## Types of templates
There are two packages operating with templates providing same interface
1. text/template- small text
1. html/template-to generate HTML output safe against code injection[xss attack] (bigger file)
## Actions
template’s syntax
```
{{ }}
```
 Templates are provided to the appropriate functions either as string or as “raw string”. Actions represents the data evaluations, functions or control loops. They’re delimited by `{{ }}` . Other, non delimited parts are left untouched.

## Data evaluations
Usually, when using templates, you’ll bind them to some data structure (e.g. struct) from which you’ll obtain data. To obtain data from a struct, you can use the `{{ .FieldName }}` action, which will replace it with `FieldName` value of given struct, on parse time. The struct is given to the `Execute` function, which we’ll cover later.

There’s also the `{{.}}` action that you can use to refer to a value of non-struct types.

## Conditions
You can also use `if loops` in templates. For example, you can check if FieldName non-empty, and if it is, print its value:
```
 {{if .FieldName}} Value of FieldName is {{ .FieldName }} {{end}}
 ```

else and else if are also supported: 
```
{{if .FieldName}} // action {{ else }} // action 2 {{ end }}
```
## Loops
Using the range action you can loop through a slice. A range actions is defined using the template
```
{{range .Member}} ... {{end}} 
```
If your slice is a non-struct type, you can refer to the value using the `{{ . }}` action. In case of structs, you can refer to the value using the `{{ .Member }}` action, as already explained.

## Functions, Pipelines and Variables
Actions have several built-in functions that’re used along with pipelines to additionally parse output. Pipelines are annotated with `|` and default behavior is sending data from left side to the function on right side.

Functions are used to escape the action’s result. There’re several functions available by default such as, `html` which returns HTML escaped output, safe against code injection or `js` which returns JavaScript escaped output.

Using the `with` action, you can define variables that’re available in that with block:
```
 {{ with $x := <^>result-of-some-action<^> }} {{ $x }} {{ end }}
 ```
## Parsing Templates
The three most important and most frequently used functions are:used for text

1. New — allocates new, undefined template,
1. Parse — parses given template string and return parsed template
1. Execute — applies parsed template to the data structure and writes result to the given writer.
The following code shows above-mentioned functions in the action:
for html template use `Parse `and `Execute`

# Verifying Templates
template packages provide the `Must` functions, used to verify that a template is valid during parsing. The Must function provides the same result as if we manually checked for the error, like in the previous example.

This approach saves you typing, but if you encounter an error, your application will panic. For advanced error handling, it’s easier to use above solution instead of Must function.

The Must function takes a template and error as arguments. It’s common to provide New function as an argument to it:
```
t := template.Must(template.New("todos").Parse("You have task named \"{{ .Name}}\" with description: \"{{ .Description}}\""))
```
