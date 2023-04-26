pipeline {
    stage('Build') {
        steps {
            nohup go run main.go
        }
    }
    stage('Test') {
        steps {
            echo 'Unit testing ...'
        }
    }
}