// Code generated by goa v3.1.1, DO NOT EDIT.
//
// guideme HTTP client CLI support package
//
// Command:
// $ goa gen guide.me/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	organizationc "guide.me/gen/http/organization/client"
	stepc "guide.me/gen/http/step/client"
	walkthroughc "guide.me/gen/http/walkthrough/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `organization (list|show|add|remove|update)
step (list|add|remove)
walkthrough (list|show|add|remove|update|rename|publish)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` organization list` + "\n" +
		os.Args[0] + ` step list --id "Quaerat aut."` + "\n" +
		os.Args[0] + ` walkthrough list --id "Asperiores maxime."` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		organizationFlags = flag.NewFlagSet("organization", flag.ContinueOnError)

		organizationListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		organizationShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		organizationShowIDFlag   = organizationShowFlags.String("id", "REQUIRED", "ID of the Organization to show")
		organizationShowViewFlag = organizationShowFlags.String("view", "", "")

		organizationAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		organizationAddBodyFlag = organizationAddFlags.String("body", "REQUIRED", "")

		organizationRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		organizationRemoveIDFlag = organizationRemoveFlags.String("id", "REQUIRED", "ID of Organization to remove")

		organizationUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		organizationUpdateBodyFlag = organizationUpdateFlags.String("body", "REQUIRED", "")

		stepFlags = flag.NewFlagSet("step", flag.ContinueOnError)

		stepListFlags  = flag.NewFlagSet("list", flag.ExitOnError)
		stepListIDFlag = stepListFlags.String("id", "REQUIRED", "ID of Walkthrough to search for steps ")

		stepAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		stepAddBodyFlag = stepAddFlags.String("body", "REQUIRED", "")

		stepRemoveFlags    = flag.NewFlagSet("remove", flag.ExitOnError)
		stepRemoveBodyFlag = stepRemoveFlags.String("body", "REQUIRED", "")

		walkthroughFlags = flag.NewFlagSet("walkthrough", flag.ContinueOnError)

		walkthroughListFlags  = flag.NewFlagSet("list", flag.ExitOnError)
		walkthroughListIDFlag = walkthroughListFlags.String("id", "REQUIRED", "ID of Organization to search for ")

		walkthroughShowFlags    = flag.NewFlagSet("show", flag.ExitOnError)
		walkthroughShowIDFlag   = walkthroughShowFlags.String("id", "REQUIRED", "ID of the Walkthrough to show")
		walkthroughShowViewFlag = walkthroughShowFlags.String("view", "", "")

		walkthroughAddFlags    = flag.NewFlagSet("add", flag.ExitOnError)
		walkthroughAddBodyFlag = walkthroughAddFlags.String("body", "REQUIRED", "")

		walkthroughRemoveFlags  = flag.NewFlagSet("remove", flag.ExitOnError)
		walkthroughRemoveIDFlag = walkthroughRemoveFlags.String("id", "REQUIRED", "ID of Walkthrough to remove")

		walkthroughUpdateFlags    = flag.NewFlagSet("update", flag.ExitOnError)
		walkthroughUpdateBodyFlag = walkthroughUpdateFlags.String("body", "REQUIRED", "")

		walkthroughRenameFlags    = flag.NewFlagSet("rename", flag.ExitOnError)
		walkthroughRenameBodyFlag = walkthroughRenameFlags.String("body", "REQUIRED", "")

		walkthroughPublishFlags  = flag.NewFlagSet("publish", flag.ExitOnError)
		walkthroughPublishIDFlag = walkthroughPublishFlags.String("id", "REQUIRED", "ID of Walkthrough to be published")
	)
	organizationFlags.Usage = organizationUsage
	organizationListFlags.Usage = organizationListUsage
	organizationShowFlags.Usage = organizationShowUsage
	organizationAddFlags.Usage = organizationAddUsage
	organizationRemoveFlags.Usage = organizationRemoveUsage
	organizationUpdateFlags.Usage = organizationUpdateUsage

	stepFlags.Usage = stepUsage
	stepListFlags.Usage = stepListUsage
	stepAddFlags.Usage = stepAddUsage
	stepRemoveFlags.Usage = stepRemoveUsage

	walkthroughFlags.Usage = walkthroughUsage
	walkthroughListFlags.Usage = walkthroughListUsage
	walkthroughShowFlags.Usage = walkthroughShowUsage
	walkthroughAddFlags.Usage = walkthroughAddUsage
	walkthroughRemoveFlags.Usage = walkthroughRemoveUsage
	walkthroughUpdateFlags.Usage = walkthroughUpdateUsage
	walkthroughRenameFlags.Usage = walkthroughRenameUsage
	walkthroughPublishFlags.Usage = walkthroughPublishUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "organization":
			svcf = organizationFlags
		case "step":
			svcf = stepFlags
		case "walkthrough":
			svcf = walkthroughFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "organization":
			switch epn {
			case "list":
				epf = organizationListFlags

			case "show":
				epf = organizationShowFlags

			case "add":
				epf = organizationAddFlags

			case "remove":
				epf = organizationRemoveFlags

			case "update":
				epf = organizationUpdateFlags

			}

		case "step":
			switch epn {
			case "list":
				epf = stepListFlags

			case "add":
				epf = stepAddFlags

			case "remove":
				epf = stepRemoveFlags

			}

		case "walkthrough":
			switch epn {
			case "list":
				epf = walkthroughListFlags

			case "show":
				epf = walkthroughShowFlags

			case "add":
				epf = walkthroughAddFlags

			case "remove":
				epf = walkthroughRemoveFlags

			case "update":
				epf = walkthroughUpdateFlags

			case "rename":
				epf = walkthroughRenameFlags

			case "publish":
				epf = walkthroughPublishFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "organization":
			c := organizationc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = organizationc.BuildShowPayload(*organizationShowIDFlag, *organizationShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = organizationc.BuildAddPayload(*organizationAddBodyFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = organizationc.BuildRemovePayload(*organizationRemoveIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = organizationc.BuildUpdatePayload(*organizationUpdateBodyFlag)
			}
		case "step":
			c := stepc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = stepc.BuildListPayload(*stepListIDFlag)
			case "add":
				endpoint = c.Add()
				data, err = stepc.BuildAddPayload(*stepAddBodyFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = stepc.BuildRemovePayload(*stepRemoveBodyFlag)
			}
		case "walkthrough":
			c := walkthroughc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = walkthroughc.BuildListPayload(*walkthroughListIDFlag)
			case "show":
				endpoint = c.Show()
				data, err = walkthroughc.BuildShowPayload(*walkthroughShowIDFlag, *walkthroughShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = walkthroughc.BuildAddPayload(*walkthroughAddBodyFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = walkthroughc.BuildRemovePayload(*walkthroughRemoveIDFlag)
			case "update":
				endpoint = c.Update()
				data, err = walkthroughc.BuildUpdatePayload(*walkthroughUpdateBodyFlag)
			case "rename":
				endpoint = c.Rename()
				data, err = walkthroughc.BuildRenamePayload(*walkthroughRenameBodyFlag)
			case "publish":
				endpoint = c.Publish()
				data, err = walkthroughc.BuildPublishPayload(*walkthroughPublishIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// organizationUsage displays the usage of the organization command and its
// subcommands.
func organizationUsage() {
	fmt.Fprintf(os.Stderr, `The Organization service makes it possible to view, add or remove Organizations.
Usage:
    %s [globalflags] organization COMMAND [flags]

COMMAND:
    list: List all stored Organizations
    show: Show Organization by ID
    add: Add new bottle and return its ID.
    remove: Remove Organization from storage
    update: Update organization with the given IDs.

Additional help:
    %s organization COMMAND --help
`, os.Args[0], os.Args[0])
}
func organizationListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization list

List all stored Organizations

Example:
    `+os.Args[0]+` organization list
`, os.Args[0])
}

func organizationShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization show -id STRING -view STRING

Show Organization by ID
    -id STRING: ID of the Organization to show
    -view STRING: 

Example:
    `+os.Args[0]+` organization show --id "Alias neque vero." --view "default"
`, os.Args[0])
}

func organizationAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization add -body JSON

Add new bottle and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` organization add --body '{
      "name": "Creating a new request in netflix!",
      "url": "http://www.google.com/"
   }'
`, os.Args[0])
}

func organizationRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization remove -id STRING

Remove Organization from storage
    -id STRING: ID of Organization to remove

Example:
    `+os.Args[0]+` organization remove --id "Est et excepturi velit."
`, os.Args[0])
}

func organizationUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization update -body JSON

Update organization with the given IDs.
    -body JSON: 

Example:
    `+os.Args[0]+` organization update --body '{
      "id": "123abc",
      "name": "Creating a new request in netflix!",
      "url": "http://www.google.com/"
   }'
`, os.Args[0])
}

// stepUsage displays the usage of the step command and its subcommands.
func stepUsage() {
	fmt.Fprintf(os.Stderr, `The Step service makes it possible to view, add, modify or remove Steps of a Walkthrough.
Usage:
    %s [globalflags] step COMMAND [flags]

COMMAND:
    list: List all stored Steps for a given walkthrough
    add: Add new Steps to walkthrough and return ID.
    remove: Remove Steps from storage

Additional help:
    %s step COMMAND --help
`, os.Args[0], os.Args[0])
}
func stepListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] step list -id STRING

List all stored Steps for a given walkthrough
    -id STRING: ID of Walkthrough to search for steps 

Example:
    `+os.Args[0]+` step list --id "Quaerat aut."
`, os.Args[0])
}

func stepAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] step add -body JSON

Add new Steps to walkthrough and return ID.
    -body JSON: 

Example:
    `+os.Args[0]+` step add --body '{
      "step": {
         "action": "next",
         "content": "This dropdown contains values from the list of status, for our scenario we want to chose \'active\'",
         "placement": "left",
         "stepNumber": 1411758952,
         "target": "Aliquam non nostrum veniam et sapiente.",
         "title": "Click here to make it work!"
      },
      "wtId": "Eos doloremque quo aut molestiae."
   }'
`, os.Args[0])
}

func stepRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] step remove -body JSON

Remove Steps from storage
    -body JSON: 

Example:
    `+os.Args[0]+` step remove --body '{
      "id": "Tempore provident at fugit libero ut recusandae.",
      "wtId": "Voluptatem est ipsa mollitia in atque."
   }'
`, os.Args[0])
}

// walkthroughUsage displays the usage of the walkthrough command and its
// subcommands.
func walkthroughUsage() {
	fmt.Fprintf(os.Stderr, `The walkthrough service makes it possible to view, add, modify or remove walkthroughs.
Usage:
    %s [globalflags] walkthrough COMMAND [flags]

COMMAND:
    list: List all stored walkthrough for a given organization
    show: Show Walkthrough by ID
    add: Add new Tutorial and return its ID.
    remove: Remove Walkthrough from storage
    update: Update Walkthrough with the given IDs.
    rename: Rename Walkthrough with the given IDs.
    publish: Publishes Walkthrough with the given IDs.

Additional help:
    %s walkthrough COMMAND --help
`, os.Args[0], os.Args[0])
}
func walkthroughListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough list -id STRING

List all stored walkthrough for a given organization
    -id STRING: ID of Organization to search for 

Example:
    `+os.Args[0]+` walkthrough list --id "Asperiores maxime."
`, os.Args[0])
}

func walkthroughShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough show -id STRING -view STRING

Show Walkthrough by ID
    -id STRING: ID of the Walkthrough to show
    -view STRING: 

Example:
    `+os.Args[0]+` walkthrough show --id "Minima itaque vitae." --view "default"
`, os.Args[0])
}

func walkthroughAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough add -body JSON

Add new Tutorial and return its ID.
    -body JSON: 

Example:
    `+os.Args[0]+` walkthrough add --body '{
      "baseURL": "http://www.google.com/",
      "name": "How to create a new process using the exception condition.",
      "organization": "Ea accusamus enim repudiandae.",
      "publishedURL": "Corporis est.",
      "status": "draft | published"
   }'
`, os.Args[0])
}

func walkthroughRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough remove -id STRING

Remove Walkthrough from storage
    -id STRING: ID of Walkthrough to remove

Example:
    `+os.Args[0]+` walkthrough remove --id "Neque sit non unde."
`, os.Args[0])
}

func walkthroughUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough update -body JSON

Update Walkthrough with the given IDs.
    -body JSON: 

Example:
    `+os.Args[0]+` walkthrough update --body '{
      "baseURL": "http://www.google.com/",
      "id": "123abc",
      "name": "How to create a new process using the exception condition.",
      "organization": "Eos earum fugiat quia assumenda odit.",
      "publishedURL": "Odit asperiores.",
      "status": "draft | published"
   }'
`, os.Args[0])
}

func walkthroughRenameUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough rename -body JSON

Rename Walkthrough with the given IDs.
    -body JSON: 

Example:
    `+os.Args[0]+` walkthrough rename --body '{
      "id": "Adipisci minima cum consequatur occaecati commodi laudantium.",
      "name": "Aut et omnis."
   }'
`, os.Args[0])
}

func walkthroughPublishUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough publish -id STRING

Publishes Walkthrough with the given IDs.
    -id STRING: ID of Walkthrough to be published

Example:
    `+os.Args[0]+` walkthrough publish --id "Veniam cupiditate quia aperiam illo enim."
`, os.Args[0])
}
