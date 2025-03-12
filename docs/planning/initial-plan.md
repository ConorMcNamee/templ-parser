# Initial Plan

The first stage is to simply be able to parse an entire project and generate all of the templ files into HTML. This should be fairly simple as the majority of this is documented on the templ docs

1. Have a set folder for templ files and get all files within that folder and subdirectories
2. Find all functions within the templ go files
3. Render all files and save as html
4. Output to templ directory for all html files

## Additional steps

Beyond simply generating the HTML there are a number of different features that we can implement. See the list below

1. VScode extension - be able to preview your changes directly in vscode. Similar to a markdown viewer. It is easy to generate html files however, having an interface to view them is the tricky part
2. Be able to automatically templ files in a project and where they are located
3. Support both manual generation and hot reloading
4. Be completely secure and be air tight and perform security audits
5. Be more performant. I imagine V1 will be slow and won't work very well so improve performance

## Final Goal

For the final project, I want to have a tool which can parse a go project with templ files, local the templ files, render the templ files and provide a debug version of the templ function with the css styling and stubbed data

# (DONE) Step 1: Get all templ files

This step is extremely easy. Simply have a set directory (this can easily be changed to a variable later), search this directory for any files with the suffix of _templ.go and return them as a list of strings

# Step 2: Find all the Go templ functions

This is a bit trickier

# Step 3: