CREATE TABLE taggedmark (
    id INTEGER PRIMARY KEY NOT NULL,
    -- title TEXT NOT NULL,
    url TEXT NOT NULL,
    UNIQUE (url)
) STRICT;
CREATE TABLE tag (
    id INTEGER PRIMARY KEY NOT NULL,
    name INTEGER NOT NULL,
    UNIQUE (name),
    CHECK (name NOT LIKE '% %') -- no spaces
) STRICT;

CREATE TABLE taggedmark_tag (
    tag_id INTEGER NOT NULL,
    taggedmark_id INTEGER NOT NULL,
    FOREIGN KEY (tag_id) REFERENCES tag(id) ON DELETE CASCADE,
    FOREIGN KEY (taggedmark_id) REFERENCES taggedmark(id) ON DELETE CASCADE,
    PRIMARY KEY (tag_id, taggedmark_id)
) STRICT;