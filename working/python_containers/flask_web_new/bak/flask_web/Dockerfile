FROM nanodatacenter/panadas

RUN pip install gunicorn
RUN pip install Flask
RUN pip install Flask-HTTPAuth
RUN pip install Flask-SocketIO
RUN pip install Werkzeug
RUN pip install Jinja2
RUN pip install flask_httpauth
#RUN pip install urllib
RUN pip install flatten-dict
ENV CONTAINER_NAME "flask_web"
COPY . .


