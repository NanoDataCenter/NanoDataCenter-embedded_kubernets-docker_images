#!/usr/bin/env python3
"""
Create and populate a minimal PostgreSQL schema for full text search
"""

import postgresql

DB_NAME = 'pgfts'
DB_HOST = 'localhost' # Uses a local socket
DB_USER = 'fts_user'

DB = postgresql.open(host=DB_HOST, database=DB_NAME, user=DB_USER)

def load_db():
    """Add sample data to the database"""

    ins = DB.prepare("""
        INSERT INTO fulltext_search (doc)
        VALUES ($1)
    """)
    ins('Sketching the trees')
    ins('Found in schema.org')
    ins('Sketched out in schema.org')
    ins('Girl on a train')

def init_db():
    """Initialize our database"""
    DB.execute("DROP TABLE IF EXISTS fulltext_search")
    DB.execute("""
        CREATE TABLE fulltext_search (id SERIAL, doc TEXT, tsv TSVECTOR)
    """)
    DB.execute("""
        CREATE TRIGGER tsvupdate BEFORE INSERT OR UPDATE ON fulltext_search
        FOR EACH ROW EXECUTE PROCEDURE
        tsvector_update_trigger(tsv, 'pg_catalog.english', doc)
    """)
    DB.execute("CREATE INDEX fts_idx ON fulltext_search USING GIN(tsv)")

if __name__ == "__main__":
    init_db()
    load_db()
pi@blockchain:~/postgresql-full-text-search-engine $

#!/usr/bin/env python3
"""
Serve up search results as JSON via REST requests

Provide JSON results including the ID and the search snippet for given search
requests.

Ultimately needs to support advanced search as well, including NOT operators
and wildcards.
"""

import json
import postgresql
import flask
import urllib.parse

app = flask.Flask(__name__)

# Port on which JSON should be served up
PORT = 8001

# Database connection info
DB_NAME = 'pgfts'
DB_HOST = 'localhost'
DB_USER = 'fts_user'

DB = postgresql.open(host=DB_HOST, database=DB_NAME, user=DB_USER)

@app.route("/search/<query>/")
@app.route("/search/<query>/<int:page>")
@app.route("/search/<query>/<int:page>/<int:limit>")
def search(query, page=0, limit=10):
    """Return JSON formatted search results, including snippets and facets"""

    query = urllib.parse.unquote(query)
    results = __get_ranked_results(query, limit, page)
    count = __get_result_count(query)

    resj = json.dumps({
        'query': query,
        'results': results,
        'meta': {
            'total': count,
            'page': page,
            'limit': limit,
            'results': len(results)
        }
    })
    return flask.Response(response=str(resj), mimetype='application/json')

def __get_ranked_results(query, limit, page):
    """Simple search for terms, with optional limit and paging"""

    sql = """
        WITH q AS (SELECT plainto_tsquery($1) AS query),
        ranked AS (
            SELECT id, doc, ts_rank(tsv, query) AS rank
            FROM fulltext_search, q
            WHERE q.query @@ tsv
            ORDER BY rank DESC
            LIMIT $2 OFFSET $3
        )
        SELECT id, ts_headline(doc, q.query, 'MaxWords=75,MinWords=25,ShortWord=3,MaxFragments=3,FragmentDelimiter="||||"')
        FROM ranked, q
        ORDER BY ranked DESC
    """

    cur = DB.prepare(sql)
    results = []
    for row in cur(query, limit, page*limit):
        results.append({
            'id': row[0],
            'snippets': row[1].split('||||')
        })

    return results

def __get_result_count(query):
    """Gather count of matching results"""

    sql = """
        SELECT COUNT(*) AS rescnt
        FROM fulltext_search
        WHERE plainto_tsquery($1) @@ tsv
    """
    cur = DB.prepare(sql)
    count = cur.first(query)
    return count

if __name__ == "__main__":
    app.debug = True
    app.run(port=PORT)
    
#!/usr/bin/env python3
"""
Simple search UI for PostgreSQL full-text search
"""

import json
import flask
import os
import urllib.parse, urllib.request
from jinja2 import Environment, FileSystemLoader

JSON_HOST = "http://localhost:8001" # Where the restserv service is running

env = Environment(loader=FileSystemLoader(searchpath="%s/templates" % os.path.dirname((os.path.realpath(__file__)))), trim_blocks=True)
app = flask.Flask(__name__)

@app.route("/")
def index():
    """Serve up the basic search page"""
    template = env.get_template('index.html')
    return(template.render())

@app.route("/search")
def search():
    """Simple search for terms, with optional limit and paging"""
    query = flask.request.args.get('query', '')
    page = flask.request.args.get('page', '')
    jsonu = u"%s/search/%s/" % (JSON_HOST, urllib.parse.quote_plus(query.encode('utf-8')))
    if page:
        jsonu = u"%s%d" % (jsonu, int(page))
    res = json.loads(urllib.request.urlopen(jsonu).read().decode('utf-8'))
    template = env.get_template('results.html')
    return(template.render(
        terms=res['query'].replace('+', ' '),
        results=res,
        request=flask.request
    ))

if __name__ == "__main__":
    app.debug = True
    print(JSON_HOST)
    app.run()
pi@blockchain:~/postgresql-full-text-search-engine/webapp $


































    
    
