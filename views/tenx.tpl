<!DOCTYPE html>

<html>
  <head>
    <title>Ceph Web</title>

    <!-- Css -->
    <link rel="stylesheet" type="text/css" href="static/bower_components/bootstrap/dist/css/bootstrap.min.css">
    <link rel="stylesheet" type="text/css" href="static/bower_components/bootstrap-material-design/dist/css/material-fullpalette.min.css">
    <link rel="stylesheet" type="text/css" href="static/bower_components/bootstrap-material-design/dist/css/material.min.css">
    <link rel="stylesheet" type="text/css" href="static/bower_components/bootstrap-material-design/dist/css/ripples.min.css">
    <link rel="stylesheet" type="text/css" href="static/bower_components/bootstrap-material-design/dist/css/roboto.min.css">
    <link rel="stylesheet" type="text/css" href="static/css/style.css">

    <!-- JavaScript -->
    <script src="static/bower_components/jquery/dist/jquery.min.js"></script>
    <script src="static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
    <script src="static/bower_components/bootstrap-material-design/dist/js/material.min.js"></script>
    <script src="static/bower_components/bootstrap-material-design/dist/js/ripples.min.js"></script>
    <script src="static/js/style.js"></script>
    <script>
        // Refer to http://fezvrasta.github.io/bootstrap-material-design/#getting-started
        $(document).ready(function() {
            // This command is used to initialize some elements and make them work properly
            $.material.init();
        });
    </script>

  </head>

  <body>

<!-- For more usage of bootstrap-material-design in https://fezvrasta.github.io/bootstrap-material-design/bootstrap-elements.html -->

    <!-- Nav bar -->
    <div class="navbar navbar-warning">
      <div class="navbar-header">
        <button type="button" class="navbar-toggle" data-toggle="collapse" data-target=".navbar-warning-collapse">
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
          <span class="icon-bar"></span>
        </button>
        <a class="navbar-brand" href="javascript:void(0)">TenxCloud Ceph Dashboard</a>
      </div>
    </div>


    <div class="container">

    <!-- Hosts card -->
      <div class="panel panel-primary">
          <div class="panel-heading">
              <span class="panel-title">Ceph Hosts</span>
              <!-- <span class="togglebutton">
                <label>
                  &nbsp; &nbsp; &nbsp;<input type="checkbox" checked="">
                </label>
              </span> -->
          </div>
          <div class="panel-body">
            <p><b>Output:</b> {{.hosts}}</p>
          </div>
      </div>

      <!-- Health card -->
      <div class="panel panel-primary">
          <div class="panel-heading">
              <span class="panel-title">Ceph Health</span>
              <!-- <span class="togglebutton">
                <label>
                  &nbsp; &nbsp; &nbsp;<input type="checkbox" checked="">
                </label>
              </span> -->
          </div>
          <div class="panel-body">
            <p><b>Overall status:</b> {{.health.OverallStatus}}</p>
            <p><b>Detail:</b>  {{.health.Detail}}</p>
            <p><b>Summary:</b>  {{.health.Summary}}</p>
          </div>
      </div>

      <div class="panel panel-primary">
          <div class="panel-heading">
              <span class="panel-title">OSD Stat</span>
              <!-- <span class="togglebutton">
                <label>
                  &nbsp; &nbsp; &nbsp;<input type="checkbox" checked="">
                </label> -->
              </span>
          </div>
          <div class="panel-body">
              <p><b>OSD Total:</b> {{.osdStatItem.NumOsds}}</p>
              <p><b>OSD Up:</b> {{.osdStatItem.NumUpOsds}}</p>
              <p><b>OSD In:</b> {{.osdStatItem.NumInOsds}}</p>
          </div>
      </div>

      <div class="panel panel-primary">
          <div class="panel-heading">
              <span class="panel-title">MON Stat</span>
              <!-- <span class="togglebutton">
                <label>
                  &nbsp; &nbsp; &nbsp;<input type="checkbox" checked="">
                </label> -->
              </span>
          </div>
          <div class="panel-body">
              <p><b>Mons:</b> {{.monStatItem.MonNames}}</p>
              <p><b>Quorums:</b> {{.monStatItem.QuorumNames}}</p>
          </div>
      </div>

      <div class="panel panel-primary">
          <div class="panel-heading">
              <span class="panel-title">Pool Info</span>
              <!-- <span class="togglebutton">
                <label>
                  &nbsp; &nbsp; &nbsp;<input type="checkbox" checked="">
                </label> -->
              </span>
          </div>
          <div class="panel-body">
              <p><b>Pools:</b> {{.poolItems}}</p>
          </div>
      </div>

      <div class="panel panel-primary">
          <div class="panel-heading">
              <span class="panel-title">PG Stat</span>
              <!-- <span class="togglebutton">
                <label>
                  &nbsp; &nbsp; &nbsp;<input type="checkbox" checked="">
                </label> -->
              </span>
          </div>
          <div class="panel-body">
              <p><b>PG Total:</b> {{.pgStatItem.NumPgs}}</p>
              <p><b>PG By State:</b> {{.pgStatItem.PgsByState}}</p>
          </div>
      </div>

      <div class="panel panel-primary">
          <div class="panel-heading">
              <span class="panel-title">Disk Usage</span>
              <!-- <span class="togglebutton">
                <label>
                  &nbsp; &nbsp; &nbsp;<input type="checkbox" checked="">
                </label>
              </span> -->
          </div>
          <div class="panel-body">
              <p><b>Total disk(KB):</b> {{.diskUsageItem.TotalKb}}</p>
              <p><b>Used disk(KB):</b> {{.diskUsageItem.TotalKbUsed}}</p>
              <p><b>Available disk(KB):</b> {{.diskUsageItem.TotalKbAvail}}</p>
          </div>
      </div>

    </div><!-- End of container -->

  </body>
</html>
