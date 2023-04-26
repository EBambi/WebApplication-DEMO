pipeline {
    agent any
    stages{
        stage('Checkout') {
            steps {
                sshagent(credentials: ['WebApplication-Demo']) {
                    git url: 'https://github.com/EBambi/WebApplication-DEMO.git', branch: 'master', credentialsId: 'WebApplication-Demo'
                }
            }
        }
        /*stage('Build') {
            steps {
                sh 'go build main.go'
            }
        }*/
    }
}