# rach
Rach is an utility to help you be organized. in the items that you have at your home.
- Every items location needs to fixed.
- Every item needs to be marked and identified.
- they need to be kept in locations as per their usage.
- Items less used should be kept fatrher away.
- Items used more often needs to be kept in a closer location.
- there should be an auditing mechanism to adudit things on a regular basis
- there should be a search mechanism where induviduals can search based on keywords.
- there should be a movement mechanishm where item can be moved from one place to another.
- there should be a finish mechanishm to indicate that a certain thing is no longer available.
- An indivudual may have more than one location to take care of.

# About the software
- this is basically categorized into 2 parts
    - client-side
        - Android module
        - website module
        - cli module
    - server-side
        - server-API acceptance.
        - integrates with Database
    - Database
        - SQLLite
        - Redis
    - iac
        - terraform to crated virtual instances.
        - ansible to install necesary modules in the instances.
        - docker files.
    - docs
        - documentations


## How we propose to achive this.
We propose to achiven this in the following way.
- All items to be stored in a holder.
    - the Holder can be a rack or 
    - a shelf
    - or even another item.
- Every Item needs to be in a holder.
- Every holder needs to have a location.
    - It can be a fixed location .. or
    - A location based on parent.
- The basic objects are these.

### Object
- An object is a global system which can be anything.
- think about this as an abstract thing but can contains or it can be contained.
- but every object will be induvudidually identifiable.
- since dimensions etc are common to holdres as well as items it is categroized as object properties.
- Attributes.
    - object id 
    - object category. 
    - type of item
    - keywords
    - dimensions.
        - l
        - b
        - h
        - w

### Item
- An item is the object we are trying to organize.
- it can be anyting like a pen, or and eatable like chips  etc.
- the following attributes.
    - parent holder id
    - type of the item.
    - keywords to identify the item.
    - id of the item .. for programs to identify them.
- so the structure is as follows.
    - item
        - objectid (the objectid of the item)
        - holderid (objectid of the holder)

### holder
- A holder can be anything that holds items.
- structure
    - objectid: objectid of the holder
    - is_fixed : boolean to indicate if the object is at a fixed location. for eg. 
        - a cupboard can be considered as fixed.
        - but a file holder is not.
    