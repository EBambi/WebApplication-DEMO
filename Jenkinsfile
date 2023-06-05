pipeline {
    agent {label 'agent-one'}
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
                sh 'ssh -i "/home/jenkins2/id_rsa" ubuntu@ec2-3-133-122-156.us-east-2.compute.amazonaws.com -t "systemctl --user stop web.service"'
                sh 'scp -i "/home/jenkins2/id_rsa" main index.html greet.html ubuntu@ec2-3-133-122-156.us-east-2.compute.amazonaws.com:/home/ubuntu/'
                sh 'ssh -i "/home/jenkins2/id_rsa" ubuntu@ec2-3-133-122-156.us-east-2.compute.amazonaws.com -t "systemctl --user start web.service"'
            }
        }
    }
}