CREATE TABLE taggedmark (
    id INTEGER PRIMARY KEY NOT NULL,
    create_time TEXT NOT NULL,
    update_time TEXT NOT NULL,
    url TEXT NOT NULL,
    UNIQUE (url)
) STRICT;
CREATE TABLE tag (
    id INTEGER PRIMARY KEY NOT NULL,
    create_time TEXT NOT NULL,
    name TEXT NOT NULL,
    update_time TEXT NOT NULL,
    UNIQUE (name),
    CHECK (name NOT LIKE '% %') -- no spaces
) STRICT;

CREATE TABLE taggedmark_tag (
    tag_id INTEGER NOT NULL,
    taggedmark_id INTEGER NOT NULL,
    update_time TEXT NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE,
    FOREIGN KEY (taggedmark_id) REFERENCES taggedmark(id) ON DELETE CASCADE,
    PRIMARY KEY (tag_id, taggedmark_id)
) STRICT;