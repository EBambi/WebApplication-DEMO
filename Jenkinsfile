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
                sh 'go build main.go'
            }
        }
        stage('Test') {
            steps {
                sh 'echo "Unit testing"'
            }
        }
    }
}