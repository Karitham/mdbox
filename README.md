# mdbox

mdbox is a markdown utilities box.

The idea is to make a simple and user friendly interface to common operations, such that it is easy to have a command line first editing experience.

## Usage

Currently implemented commands are:

- mv
  
  **Example**

  ```sh
  mdbox mv docs/asset_test.png docs/asset.png
  ```

  Will update all backlinks of any markdown files in the current directory that contains a path to this file.
  
  Relative path from other directories are currently not supported.

## Help

```sh
NAME:
   mdbox - markdown utilities

USAGE:
   mdbox [global options] command [command options] [arguments...]

COMMANDS:
   mv       move a file somewhere else, updates backlinks
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```
