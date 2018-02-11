# Gollico

A Go client for the APIs provided by the French National Library on top of the Gallica Digital Library, as documented at http://api.bnf.fr/documents-numeriques

Not sure I'll implement all of the APIs, just playing around with them atm.

## Implemented

- **Table of Contents**. Given an _ark_ identifier for a document in the digital library, retrieves its Table of Content and returns it as json. Unit tested for TEI tables of contents, not for HTML table of contents - should more or less work.

## Not implemented

- Serials Holding: retrieves the years and issues for a periodical publication
- Metadata: retrieves the metadata / bibliographic data for a document
- Pages: retrieves the pages informations for a document, e.g. at which page is there a table of content in the document?
- Occurrence: retrieves the list of occurrences of a term in an indexed document
- Text: retrieves the OCRed text of a page / document
- Image: retrieves the image of a physical page of a document in various predefined formats (thmbnail, etc.)
- IIIF (International Image Interoperability Framework): standardized API to retrieve and manipulate images from the digital library