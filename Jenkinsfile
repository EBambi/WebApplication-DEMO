pipeline {
    agent {label 'worker-1'}
    stages{
        stage('Checkout') {
            steps {
                sshagent(credentials: ['WebApplication-Demo']) {
                    git url: 'https://github.com/EBambi/WebApplication-DEMO.git', branch: 'master', credentialsId: 'WebApplication-Demo'
                }
            }
        }
        stage('Build') {
            steps {
                sh 'go build .'
            }
        }
        stage('Deploy') {
            steps {
                sh 'scp -i "/home/jenkins2/id_rsa" main index.html greet.html ubuntu@ec2-18-117-92-64.us-east-2.compute.amazonaws.com:/home/ubuntu/'
                sh 'ssh -i "/home/jenkins2/id_rsa" sudo systemctl start web_demo_app.service"'
            }
        }
    }
}