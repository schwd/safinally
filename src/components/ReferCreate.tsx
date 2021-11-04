import { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import {
  makeStyles,
  Theme,
  createStyles,
  alpha,
} from "@material-ui/core/styles";
import Button from "@material-ui/core/Button";

import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import Snackbar from "@material-ui/core/Snackbar";
import Select from "@material-ui/core/Select";
import MuiAlert, { AlertProps } from "@material-ui/lab/Alert";
import React from 'react';
import TextField from '@material-ui/core/TextField';
import { ReferInterface } from "../models/IRefer";
import { DoctorInterface } from "../models/IDoctor";
import { HospitalInterface } from "../models/IHospital";
import { MedicalRecordInterface } from "../models/IMedicalRecord";
import { DiseasesInterface } from "../models/IDiseases";
import {
  MuiPickersUtilsProvider,
  KeyboardDateTimePicker,
} from "@material-ui/pickers";
import DateFnsUtils from "@date-io/date-fns";


const Alert = (props: AlertProps) => {
  return <MuiAlert elevation={6} variant="filled" {...props} />;
};

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      flexGrow: 1,
    },
    container: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      color: theme.palette.text.secondary,
    },
    text:{
        color: '#48D1CC',
        textAlign: 'center',
      },
      textdiag:{
        
      },
      combobox: {
        '& .MuiTextField-root': {
          margin: theme.spacing(2),
          width: '50ch',
        },
      },
      datetime: {

        margin: theme.spacing(2),
        width: 500,
      },
  
      textbox: {
        '& .MuiTextField-root': {
          margin: theme.spacing(2),
          width: '83ch',
        },
    },
  }),
);

function ReferCreate() {
  const classes = useStyles();
  const [selectedDate, setSelectedDate] = useState<Date | null>(new Date());
  const [medicalrecords, setMedicalRecords] = useState<MedicalRecordInterface[]>([]);
  const Doctor: DoctorInterface = (JSON.parse(localStorage.getItem("Doctor")|| ""));
  const [hospitals, setHospitals] = useState<HospitalInterface[]>([]);
  const [disease, setDiseases] = useState<DiseasesInterface[]>([]);
  const [refer, setRefer] = useState<Partial<ReferInterface>>(
    {}
  );

  const [value, setValue] = React.useState('Controlled');
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState(false);

  const apiUrl = "http://localhost:8080";
  const requestOptions = {
    method: "GET",
    headers: { Authorization: `Bearer ${localStorage.getItem("token")}`,
    "Content-Type": "application/json", },
  };



  const handleClose = (event?: React.SyntheticEvent, reason?: string) => {
    if (reason === "clickaway") {
      return;
    }
    setSuccess(false);
    setError(false);
  };

  const handleChange = (
    event: React.ChangeEvent<{ name?: string; value: unknown }>
  ) => {
    const name = event.target.name as keyof typeof refer;
    setRefer({
      ...refer,
      [name]: event.target.value,
    });
  };

  const handleDateChange = (date: Date | null) => {
    console.log(date);
    setSelectedDate(date);
  };


  const getHospitals = async () => {
    const requestOptions = {
      method: "GET",
      headers: { 
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",},
    }
    fetch(`${apiUrl}/api/ListHospitals`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data)
        if (res.data) {
          setHospitals(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getMedicalRecords = async () => {
    const requestOptions = {
      method: "GET",
      headers: { 
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",},
    }
    fetch(`${apiUrl}/api/ListMedicalRecord`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data)
        if (res.data) {
          setMedicalRecords(res.data);
        } else {
          console.log("else");
        }
      });
  };

  const getDisease = async () => {
    const requestOptions = {
      method: "GET",
      headers: { 
        Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",},
    }
    fetch(`${apiUrl}/api/ListDiseases`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data)
        if (res.data) {
          setDiseases(res.data);
        } else {
          console.log("else");
        }
      });
  };


  const convertType = (data: string | number | undefined) => {
    let val = typeof data === "string" ? parseInt(data) : data;
    return val;
  };

  const handleInputChange = (

    event: React.ChangeEvent<{ id?: string; value: any }>
 
  ) => {
 
    const id = event.target.id as keyof typeof ReferCreate;
 
    const { value } = event.target;
 
    setRefer({ ...refer, [id]: value });
 
  };

function submit() {
  let data = {
    DoctorID: convertType(Doctor.ID),
    MedicalRecordID: convertType(refer.MedicalRecordID),
    HospitalID: convertType(refer.HospitalID),
    Cause: refer.Cause ?? "",
    Date: selectedDate,
    DiseaseID: convertType(refer.DiseaseID),
  };

  const requestOptionsPost = {
    method: "POST",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",},
    body: JSON.stringify(data),
  };

  fetch(`${apiUrl}/api/CreateRefer`, requestOptionsPost)
    .then((response) => response.json())
    .then((res) => {
      if (res.data) {
        setSuccess(true);
      } else {
        setError(true);
      }
    });
}

  
  useEffect(() => {
    getMedicalRecords();
    getHospitals();
    getDisease();
  }, []);



return (

    <div>
        <Snackbar open={success} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="success">
          บันทึกข้อมูลสำเร็จ
        </Alert>
      </Snackbar>
      <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
        <Alert onClose={handleClose} severity="error">
          บันทึกข้อมูลไม่สำเร็จ
        </Alert>
      </Snackbar>
      <br/><br/> 
      <Typography variant="h5" className={classes.text}>
            แบบบันทึกการส่งต่อผู้ป่วยเกินศักยภาพ
          </Typography>
  
      <br/>
      <Grid className={classes.paper} container spacing={0}>
        
      <Grid item xs={3}></Grid>
        <Grid item xs={3}>
         แพทย์
         <form className={classes.combobox} noValidate >
          
         <Select
                native
                defaultValue={0}
                disabled
                variant="outlined"
              >
                  <option value={0}>
                    {Doctor.Name}
                  </option>
                
              </Select>
    
         </form>
        </Grid>

        
        <Grid item xs={3} >
          ผู้ป่วย
          <form className={classes.combobox} noValidate>
          
          <Select
         
                native
                autoWidth
                variant="outlined"
                value={refer.MedicalRecordID}
                onChange={handleChange}
                inputProps={{
                  name: "MedicalRecordID",
                }}
              >
                <option aria-label="None" value="">
                  กรุณาเลือกผู้ป่วย
                </option>
                {medicalrecords.map((item: MedicalRecordInterface) => (
                  <option value={item.ID} key={item.ID}>
                    {item.Patient_Name}
                  </option>
                ))}
              </Select>
   
        </form>
           
        
        </Grid>
        <Grid item xs={3}></Grid>
        
        <Grid item xs={3}></Grid>  
        <Grid item xs={3}>
         โรงพยาบาล
         <form className={classes.combobox} noValidate>
          
         <Select
         
         native
         variant="outlined"
         value={refer.HospitalID}
         onChange={handleChange}
         inputProps={{
           name: "HospitalID",
         }}
       >
         <option aria-label="None" value="">
           กรุณาเลือกโรงพยาบาล
         </option>
         {hospitals.map((item: HospitalInterface) => (
           <option value={item.ID} key={item.ID}>
             {item.Name}
           </option>
         ))}
       </Select>
    
         </form>
        </Grid>
 
        <Grid item xs={3}>
         โรค
         <form className={classes.combobox} noValidate>
          
         <Select
         
         native
         variant="outlined"
         value={refer.DiseaseID}
         onChange={handleChange}
         inputProps={{
           name: "DiseaseID",
         }}
       >
         <option aria-label="None" value="">
           กรุณาเลือกโรค
         </option>
         {disease.map((item: DiseasesInterface) => (
           <option value={item.ID} key={item.ID}>
             {item.Name}
           </option>
         ))}
       </Select>
    
         </form>
        </Grid>

        
        <Grid item xs={3}></Grid> 
        
        <Grid item xs={3}></Grid> 
        <Grid item xs={3} className={classes.textdiag}>
          สาเหตุ
        <div>
        <form className={classes.textbox} noValidate>
      <div>
        <TextField
          id="Cause"
          multiline
          rows={2}
          variant="outlined"
          value={refer.Cause || ""}
          onChange={handleInputChange}
        />
      </div>
    </form>
        </div>
        </Grid>
        <Grid item xs={6}></Grid>

        <Grid item xs={3}></Grid> 
    
        <Grid item xs={3}>
          วันที่
          <form className={classes.combobox} noValidate>
      

      <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDateTimePicker
                  id="Date"
                  name="Date"
                  value={selectedDate}
                  onChange={handleDateChange}
                  label=""
                  minDate={new Date("2018-01-01T00:00")}
                  format="yyyy/MM/dd hh:mm a"
                />
              </MuiPickersUtilsProvider>
    </form>
        
        </Grid>
        <Grid item xs={7}>

        </Grid>
        <Grid item xs={3}>
        <Button/>
        </Grid>
      </Grid>
      <Grid item xs={12}>
            <Button
              component={RouterLink}
              to="/"
              variant="contained"
            >
              กลับ
            </Button>
            <Button
              style={{ float: "right" }}
              variant="contained"
              onClick={submit}
              color="primary"
            >
              บันทึก
            </Button>
          </Grid>
    </div>
  );
}

export default ReferCreate;