import os
import tempfile
import time
from subprocess import call

###############################################################################
# Create note class
###############################################################################
class Note():
    NOTEDIR = '/Users/fgimenez/Dropbox/notes'
    TAGID = '#TAGS'
    EDITOR = os.environ.get('EDITOR', 'vim')

    ############################################################################
    # CONSTRUCTORS
    ############################################################################
    # Generate a note programmatically
    def __init__(self,
            note_text = None,
            tag_string = None,
            note_name = None):

        # TEXT
        if not note_text:
            self.text = ''
        else:
            self.text = note_text

        # TAGS
        if not tag_string:
            self.tags = set()
        else:
            self.tags = Note.parse_tag_string( tag_string )

        # NAME
        if not note_name:
            note_name = time.strftime('%Y-%m-%d_%H-%M-%S')

        # PATH
        if os.path.isabs(note_name):
            self.note_path = note_name
        else:
            self.note_path = os.path.join(Note.NOTEDIR,note_name)


    # Generate a note from a file
    @classmethod
    def from_file(cls,filename):

        # get file path
        if not os.path.isabs(filename):
            path = os.path.join(Note.NOTEDIR,note_name)
        else:
            path = filename

        # parse file text
        with open(path) as note_file:
            text,tag_string = Note.parse_note( note_file.read() )

        # create note object
        note = cls(
                note_text = text,
                note_tags = tag_string,
                note_name = filename)

        return note

    ############################################################################
    # STATIC METHODS
    ############################################################################
    # read text of note file and return text/tag_string
    @staticmethod
    def parse_note(s):
        text = ''
        tag_string = ''
        for line in s.splitlines():
            if line.startswith(Note.TAGID):
                tag_string = line
            else:
                text += line + os.linesep
        return( (text,tag_string) )

    # get a tag set from a tag string
    @staticmethod
    def parse_tag_string(tag_string):

        if tag_string.startswith(Note.TAGID):
            tag_string = tag_string[len(Note.TAGID):].strip() #remove beginning TAGID

        tags = set( tag_string.split(',') )
        return tags

    ############################################################################
    # CLASS METHODS
    ############################################################################
    # Generate a tag string from the tag set
    def tags2string(self):
        tag_string = '{0} {1}'.format(
            Note.TAGID,
            ','.join(self.tags) )
        return tag_string

    # Get the note text from an editor call
    def note_text_from_editor(self):
        tmp = tempfile.NamedTemporaryFile(delete=False)
        tmp_name = tmp.name
        tmp.write('\n' + self.tags2string())
        tmp.flush()
        call([Note.EDITOR,tmp.name])
        tmp.close()

        ## reopen file read after editor finishes changes
        tmp = open(tmp_name)
        self.text,tag_string = Note.parse_note( tmp.read() )
        self.tags.update( Note.parse_tag_string(tag_string) )

        # close and delete tmpfile
        tmp.close()
        os.remove(tmp_name)

    def write_note_file(self):
        if not os.path.exists(Note.NOTEDIR):
            os.mkdir(Note.NOTEDIR)

        with open(self.note_path,'w') as note_file:
            note_file.write(self.text)
            note_file.write(self.tags2string())
