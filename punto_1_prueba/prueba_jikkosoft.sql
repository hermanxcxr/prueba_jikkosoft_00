--tabla lenguaje (lenguaje predeterminado en que le aparecerán los contenidos de la página al usuario)
CREATE TABLE languages (
	id serial PRIMARY KEY,
	user_language char(2) NOT NULL UNIQUE
);

INSERT INTO languages(user_language) VALUES ('ES'),('EN'),('PT');

CREATE TABLE countries (
	id serial PRIMARY KEY,
	user_country char(2) NOT NULL UNIQUE
);

INSERT INTO countries(user_country) VALUES ('CO'),('EC'),('VE');

--tabla usuarios 
CREATE TABLE users (
  id BIGSERIAL PRIMARY KEY,
  username VARCHAR(50) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  password_hash VARCHAR(255) NOT NULL,  
  first_name VARCHAR(100),
  last_name VARCHAR(100),
  language_id integer REFERENCES languages(id),
  country_id integer REFERENCES countries(id),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);



--tabla publicaciones
CREATE TABLE posts (
  id BIGSERIAL PRIMARY KEY,
  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE, --en caso que sea eliminado el usuario, elimina sus posts
  title VARCHAR(300) NOT NULL,
  slug VARCHAR(300) NOT NULL UNIQUE,
  content TEXT NOT NULL,
  is_published BOOLEAN NOT NULL DEFAULT FALSE,
  published_at TIMESTAMP WITH TIME ZONE,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

CREATE INDEX idx_posts_users_id ON posts(user_id);
CREATE INDEX idx_posts_published_at ON posts(published_at) WHERE is_published = true;


--tabla comentarios
CREATE TABLE comments (
  id BIGSERIAL PRIMARY KEY,
  post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE, -- en caso que se elimine el post, se eliminan todos los comentarios
  user_id BIGINT REFERENCES users(id) ON DELETE SET NULL, -- en caso que se elimine el usuario, el post se mantiene pero cambia a null
  content TEXT NOT NULL,
  parent_comment_id BIGINT REFERENCES comments(id) ON DELETE CASCADE, --en caso de que se elimine su comentario padre, elimina los comentarios sub-secuentes
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT now() NOT NULL
);

CREATE INDEX idx_comments_posts ON comments(posts_id);
CREATE INDEX idx_comments_users ON comments(users_id);


--tabla etiquetas (se usa para categorizar la publicación)
CREATE TABLE tags (
  id BIGSERIAL PRIMARY KEY,
  name VARCHAR(100) NOT NULL UNIQUE
);

-- tabla de integración de etiquetas y posts
CREATE TABLE post_tags (
  post_id BIGINT NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  tag_id BIGINT NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
  PRIMARY KEY (post_id, tag_id)
);



-- función de actualización
CREATE OR REPLACE FUNCTION trg_set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = now();
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trg_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW EXECUTE FUNCTION trg_set_updated_at();

CREATE TRIGGER trg_posts_updated_at
BEFORE UPDATE ON posts
FOR EACH ROW EXECUTE FUNCTION trg_set_updated_at();

CREATE TRIGGER trg_comments_updated_at
BEFORE UPDATE ON comments
FOR EACH ROW EXECUTE FUNCTION trg_set_updated_at();








