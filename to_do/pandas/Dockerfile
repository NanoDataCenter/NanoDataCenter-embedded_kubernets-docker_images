FROM nanodatacenter/basic_python_container
ENV CONTAINER_NAME "pandas test"
#RUN apt-get install -qq -y python-numpy python-scipy
#RUN apt-get install -qq -y python-pandas python-sympy python-nose
RUN pip install pandas
RUN apt-get install gfortran libatlas-base-dev libopenblas-dev liblapack-dev -y
RUN pip install  scikit-learn   --index-url https://piwheels.org/simple 
COPY . .




