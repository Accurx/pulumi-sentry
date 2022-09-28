using System.Collections.Generic;
using System.Threading.Tasks;
using Pulumi;
using Pulumi.Sentry;

namespace Sentry
{
   class Sentry : Stack
   {
      static Task<int> Main() => Deployment.RunAsync<Sentry>();
      
      [Output] public Output<string> ClientKeyDsn { get; set; }
      
      public Sentry()
      {
         var project = new Project("test-project", new ProjectArgs
         {
            Organization = "test-org",
            Team = "test-team",
            Platform = "dotnet",
         });

         var clientKey = new Key("test-key", new KeyArgs
         {
            Organization = "test-org",
            Project = project.Name
         },
            new CustomResourceOptions{ AdditionalSecretOutputs = new List<string>{"dsnPublic"}});

         ClientKeyDsn = clientKey.DsnPublic;
      }
   }
}
