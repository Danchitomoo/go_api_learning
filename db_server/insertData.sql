insert into articles (title, contents, username, nice, created_at) values
    ('firstPost', 'This is my first blog', 'kuri', 2, now());
insert into articles (title, contents, username, nice, created_at) values
    ('2ndPost', 'Second blog', 'kuro', 4, now());

insert into comments (article_id, message, created_at) values
    (1, '1st comment yeah', now());
insert into comments (article_id, message, created_at) values
    (1, 'Great', now());