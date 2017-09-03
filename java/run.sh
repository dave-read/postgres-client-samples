
export CLASSPATH="$PWD:$PWD/lib/postgresql-42.1.4.jar"
echo "Compiling ..."
rm -f ReadTable.class
javac ReadTable.java
echo "Running ..."
java ReadTable

