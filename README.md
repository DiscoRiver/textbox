### Textbox
The textbox package is a set of tools for building complex text adventures.

### About
A lot of inspiration was taken from [this](https://gist.github.com/4poc/c279d24157af06d12cbe) gist, however extensive modifications are being made to facilitate a more complex game world. 

Some of the key differences will be;

- Simple management of the map, it's rooms, and their contents. 
- Multiple maps (For example, an adventure could take place in a high-rise office building, and you could have a different map for each floor)
- Three dimensional travel
- Reusable objects

### Cartography Reference
Currently, rooms are organised to form a map, and each room has a "Position". A basic map might look like this;

|   |   |   |   |   |
|---|---|---|---|---|
| 1-3 | 2-3 | 3-3 | 4-3 | 5-3 |
| 1-2 | 2-2 | 3-2 | 4-2 | 5-2 |
| 1-1 | 2-1 | 3-1 | 4-1 | 5-1 |

So, in this case the map might be a house, so the let's say the player would enter into grid "3-1", then "MOVE NORTH" would take them to 3-2, etc.

It's important that the grid is organised in this pattern, and expanded as required. Not all tiles need to be present, for example, yours could look like this;

|   |   |   |   |   |
|---|---|---|---|---|
| 1-3 | 2-3 | 3-3 | 4-3 | 5-3 |
| 1-2 |     | 3-2 | 4-2 | 5-2 |
|     |     | 3-1 | 4-1 | 5-1 |

This allows you to build a comlex map, but it's important that each position retains it's identifier, because this affects movement. 

In the above example, grid "1-2" might be a study, and the only way to get there is to go via a corridor, which is grid "2-3", and "1-3". When moving, a calculation is made to determine the room you want to move into, and that relies on these numbers being in the correct position. 

A feature will also be added to include a z-axis, which will account for elevation. 

More docs to come soon...
