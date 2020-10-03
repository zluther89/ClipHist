# Clip Hist

A simple, lightweight clipboard text history app for Mac OS X.

## Usage

Clone the git repository to your local machine. Either use the Github UI or run

```bash
git clone https://github.com/zluther89/ClipHist.git
```

in the directory you wish to download the folder into.

Run the ClipHist binary, or if you have go installed run in the root directory

```bash
go run *.go
```

Your clipboard history will not be recorded until the binary is run. After running the command you can access your Clipboard history by opening a web browser and typing this into your URL.

```bash
localhost:3000
```

To add old items back into your clipboard, simply click on entry. This will add the item back to your clipboard and it can be pasted as necessary.

Clip Hist records the previous 100 items you have copied into your clipboard only.

### Upcoming features:

- Support for larger history and paginated results
- Seperate Tabs for items containing URL's, phone numbers and addresses.
- Organizing entries by day
- Support for searching through entries
- Support for deleting old entries
