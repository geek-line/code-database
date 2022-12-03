import AWS from 'aws-sdk'
export const albumBucketName = 'code-database-images-ver2'
const bucketRegion = 'ap-northeast-1'
const IdentityPoolId = 'ap-northeast-1:732737cb-e31c-4318-ae61-a8bf0ab12299'

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
