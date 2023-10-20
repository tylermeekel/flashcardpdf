# FlashcardPDF
Generates a PDF of flash cards from an input text file. By default uses the "with Hashtags" output format from Omnisets, but can be configured to use whatever delineator you want.

## Usage
### I/O
By default the program will read from .txt file ```./cards.txt``` to use for the input, and output the .pdf file to ```./docs/flashcards.pdf```

### Flags
```
Usage of flashcardpdf:
    -cd string
        Sets the delineator in between each card. (Can cause errors) (default "####")
    -i string
        Sets the input filepath. (default "./cards.txt")
    -o string
        Sets the output filepath. Make sure to include the ".pdf" extension (default "./docs/flashcards.pdf")
    -qd string
        Sets the delineator in between a question and answer. (Can cause errors) (default "##")
    -afs float
        Sets the font size for the answer side of the flashcard. (default 20)
    -qfs float
        Sets the font size for the question side of the flashcard. (default 25)
```
### Exporting from Omnisets
1. Click on selected StudySet from the Dashboard.
2. Click on the export button (looks like a piece of paper with an up arrow)
3. Click on the "with Hashtags" option
4. Copy the text that is shown underneath
5. Paste into a cards.txt file 