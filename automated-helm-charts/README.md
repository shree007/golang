
# Automated Helm Chart Upload to S3 and JFrog Artifactory

helm-chart
![helm-workflow](helm-workflow.png)


Feature:

1. Load charts from base directory (Done)

2. Update dependency based on Chart.yaml (Locally for now) (Done)

3. Build dependency based on Chart.yaml (Locally for now) (Done)

4. Lint (Basic Lint checking) (Done)

5. Package (generate index file) (Done)

6. Push into Jfrog (Done)

![jfrog-upload-success](jfrog-upload-success.png)


7. Push into S3

8. Use github pipeline

ToDo:

A. 2 and 3 remotely by searching repo links

B. Write testing for each function

C. Helm lint

D. Use digest, s3Url, build version in MustAdd method of index

How to execute:
1. go mod tidy
3. go run main.go