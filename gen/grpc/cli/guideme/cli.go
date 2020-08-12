// Code generated by goa v3.1.1, DO NOT EDIT.
//
// guideme gRPC client CLI support package
//
// Command:
// $ goa gen guide.me/design

package cli

import (
	"flag"
	"fmt"
	"os"

	goa "goa.design/goa/v3/pkg"
	grpc "google.golang.org/grpc"
	organizationc "guide.me/gen/grpc/organization/client"
	stepc "guide.me/gen/grpc/step/client"
	walkthroughc "guide.me/gen/grpc/walkthrough/client"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `organization (list|show|add|remove|update)
step (list|add|remove|update)
walkthrough (list|show|add|remove|update|rename|publish)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` organization list` + "\n" +
		os.Args[0] + ` step list --message '{
      "id": "Deleniti qui."
   }'` + "\n" +
		os.Args[0] + ` walkthrough list --message '{
      "id": "Consequuntur doloremque et dolor qui."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(cc *grpc.ClientConn, opts ...grpc.CallOption) (goa.Endpoint, interface{}, error) {
	var (
		organizationFlags = flag.NewFlagSet("organization", flag.ContinueOnError)

		organizationListFlags = flag.NewFlagSet("list", flag.ExitOnError)

		organizationShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		organizationShowMessageFlag = organizationShowFlags.String("message", "", "")
		organizationShowViewFlag    = organizationShowFlags.String("view", "", "")

		organizationAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		organizationAddMessageFlag = organizationAddFlags.String("message", "", "")

		organizationRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		organizationRemoveMessageFlag = organizationRemoveFlags.String("message", "", "")

		organizationUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		organizationUpdateMessageFlag = organizationUpdateFlags.String("message", "", "")

		stepFlags = flag.NewFlagSet("step", flag.ContinueOnError)

		stepListFlags       = flag.NewFlagSet("list", flag.ExitOnError)
		stepListMessageFlag = stepListFlags.String("message", "", "")

		stepAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		stepAddMessageFlag = stepAddFlags.String("message", "", "")

		stepRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		stepRemoveMessageFlag = stepRemoveFlags.String("message", "", "")

		stepUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		stepUpdateMessageFlag = stepUpdateFlags.String("message", "", "")

		walkthroughFlags = flag.NewFlagSet("walkthrough", flag.ContinueOnError)

		walkthroughListFlags       = flag.NewFlagSet("list", flag.ExitOnError)
		walkthroughListMessageFlag = walkthroughListFlags.String("message", "", "")

		walkthroughShowFlags       = flag.NewFlagSet("show", flag.ExitOnError)
		walkthroughShowMessageFlag = walkthroughShowFlags.String("message", "", "")
		walkthroughShowViewFlag    = walkthroughShowFlags.String("view", "", "")

		walkthroughAddFlags       = flag.NewFlagSet("add", flag.ExitOnError)
		walkthroughAddMessageFlag = walkthroughAddFlags.String("message", "", "")

		walkthroughRemoveFlags       = flag.NewFlagSet("remove", flag.ExitOnError)
		walkthroughRemoveMessageFlag = walkthroughRemoveFlags.String("message", "", "")

		walkthroughUpdateFlags       = flag.NewFlagSet("update", flag.ExitOnError)
		walkthroughUpdateMessageFlag = walkthroughUpdateFlags.String("message", "", "")

		walkthroughRenameFlags       = flag.NewFlagSet("rename", flag.ExitOnError)
		walkthroughRenameMessageFlag = walkthroughRenameFlags.String("message", "", "")

		walkthroughPublishFlags       = flag.NewFlagSet("publish", flag.ExitOnError)
		walkthroughPublishMessageFlag = walkthroughPublishFlags.String("message", "", "")
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
	stepUpdateFlags.Usage = stepUpdateUsage

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

			case "update":
				epf = stepUpdateFlags

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
			c := organizationc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data = nil
			case "show":
				endpoint = c.Show()
				data, err = organizationc.BuildShowPayload(*organizationShowMessageFlag, *organizationShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = organizationc.BuildAddPayload(*organizationAddMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = organizationc.BuildRemovePayload(*organizationRemoveMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = organizationc.BuildUpdatePayload(*organizationUpdateMessageFlag)
			}
		case "step":
			c := stepc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = stepc.BuildListPayload(*stepListMessageFlag)
			case "add":
				endpoint = c.Add()
				data, err = stepc.BuildAddPayload(*stepAddMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = stepc.BuildRemovePayload(*stepRemoveMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = stepc.BuildUpdatePayload(*stepUpdateMessageFlag)
			}
		case "walkthrough":
			c := walkthroughc.NewClient(cc, opts...)
			switch epn {
			case "list":
				endpoint = c.List()
				data, err = walkthroughc.BuildListPayload(*walkthroughListMessageFlag)
			case "show":
				endpoint = c.Show()
				data, err = walkthroughc.BuildShowPayload(*walkthroughShowMessageFlag, *walkthroughShowViewFlag)
			case "add":
				endpoint = c.Add()
				data, err = walkthroughc.BuildAddPayload(*walkthroughAddMessageFlag)
			case "remove":
				endpoint = c.Remove()
				data, err = walkthroughc.BuildRemovePayload(*walkthroughRemoveMessageFlag)
			case "update":
				endpoint = c.Update()
				data, err = walkthroughc.BuildUpdatePayload(*walkthroughUpdateMessageFlag)
			case "rename":
				endpoint = c.Rename()
				data, err = walkthroughc.BuildRenamePayload(*walkthroughRenameMessageFlag)
			case "publish":
				endpoint = c.Publish()
				data, err = walkthroughc.BuildPublishPayload(*walkthroughPublishMessageFlag)
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
	fmt.Fprintf(os.Stderr, `%s [flags] organization show -message JSON -view STRING

Show Organization by ID
    -message JSON: 
    -view STRING: 

Example:
    `+os.Args[0]+` organization show --message '{
      "id": "Eos praesentium ad excepturi."
   }' --view "default"
`, os.Args[0])
}

func organizationAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization add -message JSON

Add new bottle and return its ID.
    -message JSON: 

Example:
    `+os.Args[0]+` organization add --message '{
      "name": "Creating a new request in netflix!",
      "url": "http://www.google.com/"
   }'
`, os.Args[0])
}

func organizationRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization remove -message JSON

Remove Organization from storage
    -message JSON: 

Example:
    `+os.Args[0]+` organization remove --message '{
      "id": "Fugiat et delectus quo quo animi illum."
   }'
`, os.Args[0])
}

func organizationUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] organization update -message JSON

Update organization with the given IDs.
    -message JSON: 

Example:
    `+os.Args[0]+` organization update --message '{
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
    update: Update Step with the given ID.

Additional help:
    %s step COMMAND --help
`, os.Args[0], os.Args[0])
}
func stepListUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] step list -message JSON

List all stored Steps for a given walkthrough
    -message JSON: 

Example:
    `+os.Args[0]+` step list --message '{
      "id": "Deleniti qui."
   }'
`, os.Args[0])
}

func stepAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] step add -message JSON

Add new Steps to walkthrough and return ID.
    -message JSON: 

Example:
    `+os.Args[0]+` step add --message '{
      "step": {
         "action": "click",
         "content": "This dropdown contains values from the list of status, for our scenario we want to chose \'active\'",
         "placement": "left",
         "stepNumber": 588513995,
         "target": "Ullam minima ut et aspernatur impedit iste.",
         "title": "Click here to make it work!"
      },
      "wtId": "Id vero iste voluptas."
   }'
`, os.Args[0])
}

func stepRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] step remove -message JSON

Remove Steps from storage
    -message JSON: 

Example:
    `+os.Args[0]+` step remove --message '{
      "id": "Rerum harum.",
      "wtId": "Dolor incidunt."
   }'
`, os.Args[0])
}

func stepUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] step update -message JSON

Update Step with the given ID.
    -message JSON: 

Example:
    `+os.Args[0]+` step update --message '{
      "steps": [
         {
            "action": "next",
            "content": "This dropdown contains values from the list of status, for our scenario we want to chose \'active\'",
            "id": "Sed et sit dolor aut voluptas.",
            "placement": "top",
            "stepNumber": 68608253,
            "target": "Impedit voluptatem dolor et voluptatem.",
            "title": "Click here to make it work!"
         },
         {
            "action": "next",
            "content": "This dropdown contains values from the list of status, for our scenario we want to chose \'active\'",
            "id": "Sed et sit dolor aut voluptas.",
            "placement": "top",
            "stepNumber": 68608253,
            "target": "Impedit voluptatem dolor et voluptatem.",
            "title": "Click here to make it work!"
         },
         {
            "action": "next",
            "content": "This dropdown contains values from the list of status, for our scenario we want to chose \'active\'",
            "id": "Sed et sit dolor aut voluptas.",
            "placement": "top",
            "stepNumber": 68608253,
            "target": "Impedit voluptatem dolor et voluptatem.",
            "title": "Click here to make it work!"
         },
         {
            "action": "next",
            "content": "This dropdown contains values from the list of status, for our scenario we want to chose \'active\'",
            "id": "Sed et sit dolor aut voluptas.",
            "placement": "top",
            "stepNumber": 68608253,
            "target": "Impedit voluptatem dolor et voluptatem.",
            "title": "Click here to make it work!"
         }
      ],
      "wtId": "123abc"
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
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough list -message JSON

List all stored walkthrough for a given organization
    -message JSON: 

Example:
    `+os.Args[0]+` walkthrough list --message '{
      "id": "Consequuntur doloremque et dolor qui."
   }'
`, os.Args[0])
}

func walkthroughShowUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough show -message JSON -view STRING

Show Walkthrough by ID
    -message JSON: 
    -view STRING: 

Example:
    `+os.Args[0]+` walkthrough show --message '{
      "id": "Assumenda quis laudantium."
   }' --view "default"
`, os.Args[0])
}

func walkthroughAddUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough add -message JSON

Add new Tutorial and return its ID.
    -message JSON: 

Example:
    `+os.Args[0]+` walkthrough add --message '{
      "baseURL": "http://www.google.com/",
      "name": "How to create a new process using the exception condition.",
      "organization": "Deserunt ad quo quod voluptatem beatae.",
      "publishedURL": "Aut corrupti amet in et.",
      "status": "draft | published"
   }'
`, os.Args[0])
}

func walkthroughRemoveUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough remove -message JSON

Remove Walkthrough from storage
    -message JSON: 

Example:
    `+os.Args[0]+` walkthrough remove --message '{
      "id": "Porro quis error in quia."
   }'
`, os.Args[0])
}

func walkthroughUpdateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough update -message JSON

Update Walkthrough with the given IDs.
    -message JSON: 

Example:
    `+os.Args[0]+` walkthrough update --message '{
      "baseURL": "http://www.google.com/",
      "id": "123abc",
      "name": "How to create a new process using the exception condition.",
      "organization": "Odio quia temporibus optio quasi eum aut.",
      "publishedURL": "Aliquam eligendi quis aut eum illo rem.",
      "status": "draft | published"
   }'
`, os.Args[0])
}

func walkthroughRenameUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough rename -message JSON

Rename Walkthrough with the given IDs.
    -message JSON: 

Example:
    `+os.Args[0]+` walkthrough rename --message '{
      "id": "Eum alias adipisci iste autem.",
      "name": "Id voluptatibus autem."
   }'
`, os.Args[0])
}

func walkthroughPublishUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] walkthrough publish -message JSON

Publishes Walkthrough with the given IDs.
    -message JSON: 

Example:
    `+os.Args[0]+` walkthrough publish --message '{
      "id": "Ipsum vitae et ut libero animi."
   }'
`, os.Args[0])
}
