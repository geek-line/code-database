import AWS from 'aws-sdk'
export const albumBucketName = 'code-database-images-2024'
const bucketRegion = 'ap-northeast-1'
const IdentityPoolId = 'ap-northeast-1:495e5b92-be6c-4dd2-af98-339b4ebe30af'

AWS.config.update({
  region: bucketRegion,
  credentials: new AWS.CognitoIdentityCredentials({
    IdentityPoolId: IdentityPoolId,
  }),
})
export const s3 = new AWS.S3({
  apiVersion: '2006-03-01',
  params: { Bucket: albumBucketName },
})
