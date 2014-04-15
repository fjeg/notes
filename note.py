#! /usr/bin/env python

"""
note

Usage:
    note [-h] [-m TEXT] [-t TAGS] [-n NAME] 

    -h,--help                   : Show this help message
    -m TEXT, --message TEXT     : Include note text, if not included will open and editor
    -t TAGS, --tags TAGS        : include comma separated tags
    -n NAME, --name NAME        : Name of note file
"""

from docopt import docopt
from _Note import Note

def main(args):

    n = Note(
            note_text = args['--message'],
            tag_string = args['--tags'],
            note_name = args['--name']
            )

    if not args['--message']:
        n.note_text_from_editor()

    n.write_note_file()
    print('Wrote the note to {0}'.format(n.note_path))

if __name__=='__main__':
    args = docopt(__doc__)
    main(args)
