// START1 OMIT
using System.Net;
using System.Net.Security;
using System.Security.Cryptography.X509Certificates;

public static bool OnValidationCallback(object sender, X509Certificate cert, X509Chain chain, SslPolicyErrors errors)
{
    return true;
}
// END1 OMIT

// START2 OMIT
ServicePointManager.ServerCertificateValidationCallback = new RemoteCertificateValidationCallback(OnValidationCallback);
// END2 OMIT
