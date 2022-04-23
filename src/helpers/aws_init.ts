import AWS from 'aws-sdk'
export const albumBucketName = 'code-database-images'
const bucketRegion = 'ap-northeast-1'
const IdentityPoolId = 'ap-northeast-1:0b902eb4-467e-4141-ada6-7e41b6742ba5'

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
