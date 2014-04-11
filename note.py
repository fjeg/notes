#!/usr/bin/env python
"""Usage: note.py
"""

import os
import tempfile
import time
import docopt
from subprocess import call

###############################################################################
# Create note class
###############################################################################
class Note():
    NOTEDIR = '/Users/fgimenez/Dropbox/notes'
    TAGID = '#TAGS'
    EDITOR = os.environ.get('EDITOR', 'vim')

    def __init__(self,note_name=None):
        if not os.path.exists(Note.NOTEDIR):
            os.mkdir(Note.NOTEDIR)

        if not note_name:
            note_name = time.strftime('%Y-%m-%d_%H-%M-%S')

        self.note_path = os.path.join(Note.NOTEDIR,note_name)

        self.tags = []
    # only necessary if user didn't provide note
    def note_text_from_editor(self):
        tmp = tempfile.NamedTemporaryFile(delete=False)
        tmp_name = tmp.name
        call([Note.EDITOR,tmp.name])
        tmp.close()

        ## reopen file read after editor finishes changes
        tmp = open(tmp_name)
        self.note_text = tmp.read()

        # close and delete tmpfile
        tmp.close()
        os.remove(tmp_name)

    def write_note_file(self):
        with open(self.note_path,'w') as note_file:
            note_file.write(self.note_text)
            tag_string = ','.join(self.tags)
            note_file.write('{0} {1}'.format(Note.TAGID,tag_string))



###############################################################################
# Generate note data
###############################################################################
if __name__ == '__main__':
    n = Note()
    n.note_text_from_editor()
    n.write_note_file()
    print('Written to {}'.format(n.note_path))
