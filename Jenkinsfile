pipeline {
    agent any
    stages{
        stage('Checkout') {
            steps {
                sshagent(credentials: ['WebApplication-Demo']) {
                    git url: 'https://github.com/EBambi/WebApplication-DEMO.git'
                }
            }
        }
        stage('Build') {
            steps {
                go run main.go
            }
        }
        stage('Test') {
            steps {
                echo 'Unit testing ...'
            }
        }
    }
}