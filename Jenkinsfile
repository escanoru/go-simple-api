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

        // Just for Docker
        // stage('Login') {
        //     steps {
        //         sh 'echo ${DOCKER_CREDS_PSW} | podman login -u ${DOCKER_CREDS_USR} --password-stdin'
        //     }
        // }
        stage('Push') {
            steps {
                sh 'podman push ${DOCKER_CREDS_USR}/${CONT_NAME}:latest --creds ${DOCKER_CREDS_USR}:${DOCKER_CREDS_PSW}'
            }
        }
    }
    // post {
    //     always {
    //         sh 'docker logout'
    //     }
    // }
}

