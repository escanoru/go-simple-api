pipeline {
    agent { label 'ansible' }
    options {
        buildDiscarder(logRotator(numToKeepStr: '7'))
    }
    environment {
        DOCKER_CREDS = credentials("${params.DOCKER_CREDENTIALS}")
        CONT_NAME = "${params.APP_NAME}"
    }
    parameters {
            string(name: 'APP_NAME', defaultValue: 'go-simple-api', description: 'Name of the app to push to Dockerhub')
            string(name: 'DOCKER_CREDENTIALS', defaultValue: 'docker-push', description: 'Name of the docker username and passwoed credential on Jenkins')
    }
    stages {
        stage('Build') {
            steps {         
                sh '''
                loginctl enable-linger $USER
                podman build -t ${DOCKER_CREDS_USR}/${CONT_NAME}:latest .
                '''
            }
        }
        // stage('Login') {
        //     steps {
        //         sh 'echo ${env.DOCKER_CREDENTIALS_PSW} | docker login -u ${env.DOCKER_CREDENTIALS_USR} --password-stdin'
        //     }
        // }
        // stage('Push') {
        //     steps {
        //         sh 'docker push ${env.DOCKER_CREDENTIALS_USR}/${params.APP_NAME}:latest'
        //     }
        // }
    }
    // post {
    //     always {
    //         sh 'docker logout'
    //     }
    // }
}

